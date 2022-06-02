package database

import (
	"gorm.io/gorm"
	"log"
	"urlshortener/models/database"
)

func AutoMigrateDB(DB gorm.DB) {
	err := DB.AutoMigrate(&models.Users{}, &models.Links{}, &models.Clicks{})
	if err != nil {
		log.Fatalln("[url-shortener] error running auto migrations!" + err.Error())
	}
	log.Println("[url-shortener] successfully ran auto migrations!")

}
