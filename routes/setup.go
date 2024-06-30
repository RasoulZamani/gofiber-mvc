package routes

import (
	v1 "gofiber-mvc/routes/v1"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initializes the route groups for different API versions
func SetupRoutes(app *fiber.App) {
	// Define route groups for v1
	apiV1 := app.Group("/api/v1")
	apiV1.Route("/books", v1.BookRoutes)
	apiV1.Route("/authors", v1.AuthorRoutes)

}
