package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"log"
	"project_module/database"
	"project_module/handlers"
	"project_module/static"
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
		Views:   html.NewFileSystem(static.GetPublicFiles(), ".html"),
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Get("/users", handlers.UserList)
	v1.Post("/users", handlers.UserCreate)

	// Setup static files

	// Handle not founds
	//app.Use(handlers.NotFound)
	app.Use(filesystem.New(filesystem.Config{
		Root:         static.GetPublicFiles(),
		Browse:       true,
		Index:        "public/index.html",
		NotFoundFile: "404.html",
		MaxAge:       3600,
	}))

	log.Fatal(app.Listen(*port)) // go run cmd/main.go -port=:8000
}
