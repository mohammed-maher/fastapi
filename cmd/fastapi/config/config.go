package config

import (
	"github.com/joho/godotenv"
	"log"
)

var Config *configuration


type configuration struct {
	DB    *DatabaseConfig
	Redis *RedisConfig
	JWT   *JWTConfig
	SMTP *SmtpConfig
}

func init() {
	if err:=godotenv.Load(".env");err!=nil{
		log.Panic("env file failed to load")
	}

	Config = &configuration{
		DB: loadDatabaseConfig(),
		Redis: loadRedisConfig(),
		JWT: LoadJWTConfig(),
		SMTP: LoadSmtpConfig(),
	}
}

