package main

import (
	"fmt"
	"gofiber-mvc/db"
	"gofiber-mvc/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

// Function to list all routes
func ListRoutes(app *fiber.App) {
	appRoutes := app.Stack()

	fmt.Println("Registered Routes:")
	for _, route := range appRoutes {
		for _, entry := range route {
			if entry.Method != "" {
				fmt.Printf("%s %s\n", entry.Method, entry.Path)
			}
		}
	}
}
func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome Home!")
	})

	// Add Routes
	routes.SetupRoutes(app)

	// Initialize the database connection
	database := db.NewDB()

	// Perform database migrations
	db.MigrateDB(database)

	log.Fatal(app.Listen(":3000"))

}
