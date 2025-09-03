package database

import (
	"are-you-die/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

func Open() error {
	dsn, err := config.LoadMySQLConfig()
	if err != nil {
		return fmt.Errorf("failed to load MySQL config: %w", err)
	}
	mysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	err = mysqlDB.AutoMigrate(&User{}, &Date{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	return nil
}
