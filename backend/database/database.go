package database

import (
	"couplet/api"
	"couplet/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConfigureDB(config config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Database.String()), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, err
	}

	return db, MigrateDB(db)
}

func ConnPooling(db *gorm.DB) error {
	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return nil
}

func MigrateDB(db *gorm.DB) error {
	// TODO: Add other models to auto-migration list
	return db.AutoMigrate(api.User{})
}
