package handlers

import (
	"errors"
	"gofiber-mvc/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAuthor handle reading one author by ID from db
func GetAuthor(c *fiber.Ctx) error {

	// Extract ID from the URL parameter
	id := c.Params("id")

	// Validate that the ID is a valid integer
	authorID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid author ID format",
		})
	}

	var author models.Author
	// Fetch the author from the Database by ID
	if err := Database.First(&author, uint(authorID)).Error; err != nil {
		// Check for the specific "record not found" error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Author not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An internal server error occurred",
		})
	}
	return c.JSON(author)
}

// CreateAuthor handles the creation of a new author
func CreateAuthor(c *fiber.Ctx) error {
	author := new(models.Author)

	if err := c.BodyParser(author); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	if err := Database.Create(author).Error; err != nil {
		log.Printf("Error creating author record: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create author",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(author)
}

// GetAuthors retrieves all authors from the Database and returns them as JSON
func GetAuthors(c *fiber.Ctx) error {
	var authors []models.Author

	// Fetch all authors from the Database
	if err := Database.Find(&authors).Error; err != nil {
		log.Printf("Error retrieving authors: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve authors",
			"details": err.Error(),
		})
	}

	// Return the list of authors as a JSON response
	return c.JSON(fiber.Map{
		"authors": authors,
	})
}

// UpdateAuthor handlers updates an existing author by its ID
func UpdateAuthor(c *fiber.Ctx) error {
	authorID := c.Params("id")
	if authorID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Author ID is required",
		})
	}

	// Find the author by ID
	var author models.Author
	if err := Database.First(&author, authorID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to find author",
			"details": err.Error(),
		})
	}

	// Parse the request body to update the author
	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	// Update the author in the Database
	if err := Database.Save(&author).Error; err != nil {
		log.Printf("Error updating author: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update author",
			"details": err.Error(),
		})
	}

	// Return the updated author as a JSON response
	return c.JSON(author)
}

// DeleteAuthor deletes a author by its ID
func DeleteAuthor(c *fiber.Ctx) error {
	authorID := c.Params("id")
	if authorID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Author ID is required",
		})
	}
	var author models.Author
	// find author for this id
	if err := Database.First(&author, authorID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to find author",
			"details": err.Error(),
		})
	}
	// Delete the author from the Database
	if err := Database.Delete(&author).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete  author",
			"details": err.Error(),
		})
	}

	// Return a successful response
	return c.SendStatus(fiber.StatusNoContent)
}
