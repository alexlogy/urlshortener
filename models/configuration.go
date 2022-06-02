package models

type DbConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
	MaxIdle  int    `json:"maxIdle"`
	MaxOpen  int    `json:"maxOpen"`
}

type RedisConfig struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	DbNum int    `json:"dbNum"`
}

type Configuration struct {
	Database DbConfig
	Redis    RedisConfig
}
