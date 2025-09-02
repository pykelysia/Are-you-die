package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type config_mysql struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseIP   string `json:"database_host"`
	DatabasePort string `json:"database_port"`
	DatabaseName string `json:"database_name"`
}

func LoadMySQLConfig() (string, error) {
	file, err := os.Open("config_mysql.json")
	if err != nil {
		return "", fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var config config_mysql
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return "", fmt.Errorf("failed to decode config file: %w", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.DatabaseIP,
		config.DatabasePort,
		config.DatabaseName,
	)
	fmt.Println(dsn)
	return dsn, nil
}
