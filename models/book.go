package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title    string `json:"name"`
	AuthorID uint   `json:"author_id"`
	Author   Author `gorm:"foreignKey:AuthorID;references:ID" json:"author"`
}
