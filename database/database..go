package database

import (
	"are-you-die/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open() error {
	dsn, err := config.LoadMySQLConfig()
	if err != nil {
		return fmt.Errorf("failed to load MySQL config: %w", err)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}
