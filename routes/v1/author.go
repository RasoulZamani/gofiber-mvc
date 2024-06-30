package v1

import (
	"gofiber-mvc/handlers"

	"github.com/gofiber/fiber/v2"
)

// AuthorRoutes sets up routes
func AuthorRoutes(router fiber.Router) {
	// authors := app.Group("/authors")

	router.Get("/", handlers.GetAuthors)
	router.Get("/:id", handlers.GetAuthor)
	router.Post("/", handlers.CreateAuthor)
	router.Put("/:id", handlers.UpdateAuthor)
	router.Delete("/:id", handlers.DeleteAuthor)
}
