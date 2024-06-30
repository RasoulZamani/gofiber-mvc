package handlers

import (
	"errors"
	"gofiber-mvc/models"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetBook handle reading one book by ID from db
func GetBook(c *fiber.Ctx) error {

	// Extract ID from the URL parameter
	id := c.Params("id")

	// Validate that the ID is a valid integer
	bookID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid book ID format",
		})
	}

	var book models.Book
	// Fetch the book from the Database by ID
	if err := Database.First(&book, uint(bookID)).Error; err != nil {
		// Check for the specific "record not found" error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Book not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An internal server error occurred",
		})
	}
	return c.JSON(book)
}

// CreateBook handles the creation of a new book
func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := Database.Create(book).Error; err != nil {
		log.Printf("Error creating book record: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create book",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

// GetBooks retrieves all books from the Database and returns them as JSON
func GetBooks(c *fiber.Ctx) error {
	var books []models.Book

	// Fetch all books from the Database
	if err := Database.Find(&books).Error; err != nil {
		log.Printf("Error retrieving books: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to retrieve books",
			"details": err.Error(),
		})
	}

	// Return the list of books as a JSON response
	return c.JSON(fiber.Map{
		"books": books,
	})
}

// UpdateBook handlers updates an existing book by its ID
func UpdateBook(c *fiber.Ctx) error {
	bookID := c.Params("id")
	if bookID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Book ID is required",
		})
	}

	// Find the book by ID
	var book models.Book
	if err := Database.First(&book, bookID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to find book",
			"details": err.Error(),
		})
	}

	// Parse the request body to update the book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	// Update the book in the Database
	if err := Database.Save(&book).Error; err != nil {
		log.Printf("Error updating book: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to update book",
			"details": err.Error(),
		})
	}

	// Return the updated book as a JSON response
	return c.JSON(book)
}

// DeleteBook deletes a book by its ID
func DeleteBook(c *fiber.Ctx) error {
	bookID := c.Params("id")
	if bookID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Book ID is required",
		})
	}
	var book models.Book
	// find book for this id
	if err := Database.First(&book, bookID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to find book",
			"details": err.Error(),
		})
	}
	// Delete the book from the Database
	if err := Database.Delete(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete  book",
			"details": err.Error(),
		})
	}

	// Return a successful response
	return c.SendStatus(fiber.StatusNoContent)
}
