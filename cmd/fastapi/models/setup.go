package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mohammed-maher/fastapi/config"
	"log"
)

var DB *gorm.DB

func Setup() {
	db, err := gorm.Open(config.Config.DB.DB_Driver, config.Config.DB.GenerateDSN())
	if err != nil {
		log.Fatalf("Failed to establish database connection\n%s with DSN %s", err, config.Config.DB.GenerateDSN())
	}
	db.AutoMigrate(&User{}, &Trip{}, &Car{}, &City{}, &Order{}, &Status{}, &File{})

	DB = db
}
