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
	SMTP  *SmtpConfig
	S3    *S3Config
	SMS   *SMSConfig
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("env file failed to load")
	}

	Config = &configuration{
		DB:    loadDatabaseConfig(),
		Redis: loadRedisConfig(),
		JWT:   LoadJWTConfig(),
		SMTP:  LoadSmtpConfig(),
		S3:    LoadS3Config(),
		SMS:   LoadSMSConfig(),
	}
}
