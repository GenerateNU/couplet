// Execute business logic
package controller

import (
	"gorm.io/gorm"
)

// Executes business logic
type Controller struct {
	database *gorm.DB // the active database connection
}

// Creates a new controller to handle business logic
func NewController(database *gorm.DB) Controller {
	return Controller{
		database,
	}
}
