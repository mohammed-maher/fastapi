package config

import (
	"fmt"
	"log"
)

type DatabaseConfig struct {
	DB_Name     string
	DB_Host     string
	DB_PORT     int
	DB_User     string
	DB_Password string
	DB_Driver   string
}

//Load database config
func loadDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DB_Host:     getString("DB_HOST", "localhost"),
		DB_PORT:     getInt("DB_PORT", 5432),
		DB_Name:     getString("DB_NAME", "postgres"),
		DB_Password: getString("DB_PASS", ""),
		DB_User:     getString("DB_USER", "root"),
		DB_Driver:   getString("DB_TYPE", "postgres"),
	}
}

//Generate data source name string according to database type
func (db *DatabaseConfig) GenerateDSN() string {
	switch db.DB_Driver {
	case "postgres":
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", db.DB_Host, db.DB_PORT, db.DB_User, db.DB_Name, db.DB_Password)
	case "mysql":
		return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local\n", db.DB_User, db.DB_Password, db.DB_Host, db.DB_PORT, db.DB_Name)
	}
	log.Fatal("Unknown database driver: " + db.DB_Driver)
	return ""
}
