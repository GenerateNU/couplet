package database

import (
	"fmt"
	"log/slog"

	"github.com/DATA-DOG/go-sqlmock"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connects to a PostgreSQL database through GORM
func NewDB(host string, port uint16, username string, password string, databaseName string, logger *slog.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		host, port, username, password, databaseName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: slogGorm.New(slogGorm.WithLogger(logger),
			slogGorm.WithTraceAll(),
			slogGorm.SetLogLevel(slogGorm.DefaultLogType, slog.LevelDebug)),
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
	return db.AutoMigrate(User{}, Organization{}, Event{})
}

// Creates a new mock postgres-GORM database
func NewMockDB() *gorm.DB {
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	return db
}
