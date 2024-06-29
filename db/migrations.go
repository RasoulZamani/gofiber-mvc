package db

import (
	"log"
	"my-gofiber-mvc/models"

	"gorm.io/gorm"
)

// MigrateDB performs the database migrations
func MigrateDB(db *gorm.DB) {
	// Register all models
	models := models.RegisterModels()

	// AutoMigrate all registered models
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("failed to migrate model %T: %v", model, err)
		}
	}

	log.Println("Migrations applied successfully!")
}
