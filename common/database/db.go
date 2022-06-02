package database

import (
	"gorm.io/gorm"
	"log"
	"urlshortener/common/config"
	"urlshortener/common/database/postgres"
)

var DB *gorm.DB

func InitDatabase() {
	log.Println("[url-shortener] initiating database connection..")
	DB = postgres.Connect(config.Config)
}
