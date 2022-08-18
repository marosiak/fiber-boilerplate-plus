package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"log"
	"project_module/api"
	"project_module/database"
	"project_module/static"
	"project_module/templates"
	"project_module/views"
)

var (
	port = flag.String("port", ":8000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	database.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run cmd/main.go -prod
		Views:   html.NewFileSystem(templates.GetFiles(), ".html"),
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")
	v1.Get("/users", api.UserList)
	v1.Post("/users", api.UserCreate)

	// Create server side rendered views
	app.Get("/", views.UserListView)
	app.Post("/users", views.AddUserView)

	//app.Use(api.NotFound)
	app.Use(filesystem.New(filesystem.Config{
		Root:         static.GetFiles(),
		Browse:       true,
		NotFoundFile: "private/404.html",
		MaxAge:       3600,
	}))

	app.Use(favicon.New(favicon.Config{
		File:       "public/favicon.ico",
		FileSystem: static.GetFiles(),
	}))

	log.Fatal(app.Listen(*port)) // go run cmd/main.go -port=:8000
}
