// Execute business logic
package controller

import (
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

// Executes business logic
type Controller struct {
	database *gorm.DB     // the active database connection
	logger   *slog.Logger // event logger
}

// Creates a new controller to handle business logic
func NewController(database *gorm.DB, logger *slog.Logger) (Controller, error) {
	if database == nil {
		return Controller{}, errors.New("no database specified")
	}
	db, err := database.DB()
	if err != nil || db.Ping() != nil {
		return Controller{}, errors.New("database connection failed")
	}

	return Controller{
		database,
		logger,
	}, nil
}
