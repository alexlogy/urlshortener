package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"urlshortener/models"
)

func Connect(conf *models.Configuration) *gorm.DB {
	var host = conf.Database.Host
	var port = conf.Database.Port
	var user = conf.Database.User
	var password = conf.Database.Password
	var dbName = conf.Database.DbName
	var maxIdle = conf.Database.MaxIdle
	var maxOpen = conf.Database.MaxOpen
	
	connStr := fmt.Sprintf("host= %s user=%s password=%s dbname=%s port=%d sslmode=disable Timezone=Asia/Hong_Kong", host, user, password, dbName, port)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	// connection pool settings
	dbInstance, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	dbInstance.SetMaxIdleConns(maxIdle)
	dbInstance.SetMaxOpenConns(maxOpen)

	log.Printf("[url-shortener] successfully connected to database (%s:%d).", host, port)

	return db
}
