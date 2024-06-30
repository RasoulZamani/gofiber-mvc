package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title    string `json:"name"`      // Title of the book
	AuthorID uint   `json:"author_id"` // Foreign key for the Author
	// Author   Author `gorm:"foreignKey:AuthorID;references:ID" json:"author"` // Specifies that AuthorID references the ID of Author
}
