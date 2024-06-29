package models

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Books []Book `gorm:"foreignKey:AuthorID"`
}
