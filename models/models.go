package models

// RegisterModels returns a slice of all models for AutoMigrate
func RegisterModels() []interface{} { // slice of interface help to return different model types
	return []interface{}{
		&Author{},
		&Book{},
	}
}
