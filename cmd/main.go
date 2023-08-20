package main

import (
	"log"

	"github.com/cornevanstraten/trivia-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello again, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
