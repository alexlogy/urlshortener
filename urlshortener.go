package main

import (
	"log"
	"os"
	"urlshortener/common/config"
	"urlshortener/common/database"
	"urlshortener/router"
)

func main() {
	config.LoadConfig()     // load config
	database.InitDatabase() // initialize db connection

	if os.Args[1] == "auto-migrate" {
		database.AutoMigrateDB(*database.DB)
	} else {
		r := router.SetupRouter()
		err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		if err != nil {
			log.Fatal("[url-shortener] error starting web server!")
		}
	}
}
