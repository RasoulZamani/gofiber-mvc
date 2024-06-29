package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"gofiber-mvc/db"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome Home!")
	})

	// Initialize the database connection
	database := db.NewDB()

	// Perform database migrations
	db.MigrateDB(database)

	log.Fatal(app.Listen(":3000"))

}
