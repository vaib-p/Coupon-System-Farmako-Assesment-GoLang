package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	JWTSecret   string
	FirebaseKey string
}

var AppConfig *Config

func LoadConfig() {
	cwd, _ := os.Getwd()
	log.Println("Working directory:", cwd)
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	AppConfig = &Config{
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}
