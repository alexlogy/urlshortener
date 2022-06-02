package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"urlshortener/models"
)

var Config *models.Configuration

func LoadConfig() {
	log.Println("[url-shortener] Loading config..")
	Config = InitConfig()
}

func InitConfig() (c *models.Configuration) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	c = &models.Configuration{}

	c.Database.Host = v.GetString("database.host")
	c.Database.Port = v.GetInt("database.port")
	c.Database.User = v.GetString("database.user")
	c.Database.Password = v.GetString("database.password")
	c.Database.DbName = v.GetString("database.dbName")
	c.Database.MaxIdle = v.GetInt("database.maxIdle")
	c.Database.MaxOpen = v.GetInt("database.maxOpen")

	c.Redis.Host = v.GetString("redis.host")
	c.Redis.Port = v.GetInt("redis.port")
	c.Redis.DbNum = v.GetInt("redis.dbNum")

	log.Println("[url-shortener] config summary:")
	log.Println("----------------------------------------------")
	log.Println("Database:")
	log.Println(fmt.Sprintf("Host: %s", c.Database.Host))
	log.Println(fmt.Sprintf("Port: %s", strconv.Itoa(c.Database.Port)))
	log.Println(fmt.Sprintf("User: %s", c.Database.User))
	log.Println(fmt.Sprintf("DbName: %s", c.Database.DbName))
	log.Println(fmt.Sprintf("MaxIdle: %s", strconv.Itoa(c.Database.MaxIdle)))
	log.Println(fmt.Sprintf("MaxOpen: %s", strconv.Itoa(c.Database.MaxOpen)))
	log.Println("----------------------------------------------")
	log.Println("Redis:")
	log.Println(fmt.Sprintf("Host: %s", c.Redis.Host))
	log.Println(fmt.Sprintf("Port: %s", strconv.Itoa(c.Redis.Port)))
	log.Println(fmt.Sprintf("Db Number: %s", strconv.Itoa(c.Redis.DbNum)))
	log.Println("----------------------------------------------")
	log.Println("[url-shortener] successfully loaded config")

	return
}
