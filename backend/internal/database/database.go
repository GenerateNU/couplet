// Connects to database and defines internal models
package database

import (
	"couplet/internal/database/event"
	"couplet/internal/database/org"
	"errors"
	"fmt"
	"log/slog"
	"os/user"

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

	return db, Migrate(db)
}

// Enables connection pooling on a GORM database
func EnableConnPooling(db *gorm.DB) error {
	if db == nil {
		return errors.New("nil database specified")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	return nil
}

// Performs database migrations for defined schema if necessary
func Migrate(db *gorm.DB) error {
	if db == nil {
		return errors.New("nil database specified")
	}
	// Add new models here to ensure they are migrated on startup
	return db.AutoMigrate(user.User{}, org.Org{}, event.Event{}, org.OrgTag{}, event.EventTag{})
}

// Creates a new mock postgres-GORM database
func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, mock
}
