package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	JWTSecret string
}

func Load() (*Config, error) {

	var err error = godotenv.Load()
	if err != nil {
		log.Println("error in env file.")
		

	}



	var config *Config = &Config{

		DatabaseURL: os.Getenv("DATABASE_URL") ,
		Port: os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}

	return config, nil
}
