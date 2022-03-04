package main

import (
	"party/middleware"
	"party/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New())

	party := app.Group("/", middleware.Log)
	party.Get("/", routes.Home)
	party.Post("/login", routes.Login)
	party.Static("/assets", "./assets")

	app.Listen(":8080")

}
