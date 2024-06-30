package handlers

import (
	"gofiber-mvc/db"

	"github.com/gofiber/fiber/v2"
)

// Instance of db:
var Database = db.NewDB()

// Welcome
func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome Home!")
}
