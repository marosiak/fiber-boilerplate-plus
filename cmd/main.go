package main

import (
	"flag"
	"project_module/api"
	"project_module/database"
	"project_module/sessioncontext"
	"project_module/static"
	"project_module/templates"
	"project_module/views"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberSession "github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	log "github.com/sirupsen/logrus"
)

var (
	port      = flag.String("port", ":8000", "Port to listen on")
	prod      = flag.Bool("prod", false, "Enable prefork in Production")
	verbosity = flag.Int("verbosity", 2, "Lowest is 2, highest is 6")
)

func initStaticStorage(app *fiber.App) {
	staticsFileSystem := static.GetFiles()
	app.Use(filesystem.New(filesystem.Config{
		Root:         staticsFileSystem,
		Browse:       true,
		NotFoundFile: "private/404.html",
		MaxAge:       3600,
	}))

	app.Use(favicon.New(favicon.Config{
		File:       "public/favicon.ico",
		FileSystem: staticsFileSystem,
	}))
}

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	database.Connect()

	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run cmd/main.go -prod
		Views:   html.NewFileSystem(templates.GetFiles(), ".html"),
	})
	app.Use(fiberLogger.New())
	app.Use(recover.New())

	// Init logger, the verbosity for prod is 2, for dev it's 5 or 6
	l := log.New()
	l.SetLevel(log.Level(*verbosity))
	logger := log.NewEntry(l)

	// This can be used to pass values between different views
	// usefull for displaying validation errors
	sessionCtx := sessioncontext.New(fiberSession.New(), logger)
	userViews := views.NewUserViews(sessionCtx, logger)

	// Create server side rendered views
	app.Get("/", userViews.UserListView)
	app.Post("/users", userViews.AddUserView)
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics Page"}))

	// Create a /api/v1 endpoint
	// IMPORTANT: If your app uses API + server rendered templates, you may find issue with making requests to /user
	// even if you make api endpoint equal to /api/v1/users and server rendered endpoint equal to /users, if you specify to do request on /users it will try /api/v1/users
	// the solution or hack(?) which I have found is to register API later than server rendered views
	v1 := app.Group("/api/v1")
	v1.Get("/users", api.UserList)
	v1.Post("/users", api.UserCreate)

	initStaticStorage(app)
	log.Fatal(app.Listen(*port)) // go run cmd/main.go -port=:8000
}
