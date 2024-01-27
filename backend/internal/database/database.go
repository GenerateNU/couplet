package database

import (
	"couplet/internal/api"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connects to a PostgreSQL database through GORM
func NewDb(host string, port uint16, username string, password string, databaseName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		host, port, username, password, databaseName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, err
	}

	return db, MigrateDB(db)
}

// Enables connection pooling on a GORM database
func EnableConnPooling(db *gorm.DB) error {
	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return nil
}

// Performs database migrations for defined schema if necessary
func MigrateDB(db *gorm.DB) error {
	// TODO: Add other models to auto-migration list
	return db.AutoMigrate(api.User{})
}

// Creates a new mock postgres-GORM database
func NewMockDb() *gorm.DB {
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	return db
}
