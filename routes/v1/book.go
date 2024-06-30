package v1

import (
	"gofiber-mvc/handlers"

	"github.com/gofiber/fiber/v2"
)

// BookRoutes sets up routes
func BookRoutes(router fiber.Router) {
	// books := app.Group("/books")

	router.Get("/", handlers.GetBooks)
	router.Get("/:id", handlers.GetBook)
	router.Post("/", handlers.CreateBook)
	router.Put("/:id", handlers.UpdateBook)
	router.Delete("/:id", handlers.DeleteBook)
}
