package main

import (
	"log"

	"github.com/cornevanstraten/trivia-api/database"
	"github.com/cornevanstraten/trivia-api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	setupRoutes(app)

	app.Static("/", "./public")

	app.Use(handlers.NotFound)

	log.Fatal(app.Listen(":3000"))
}
