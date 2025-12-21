package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string

	Port      string
	GinMode   string
	JWTSecret string
	AWSRegion string
	AWSKey    string
	AWSSecret string

	// Cloudinary Config
	CloudinaryCloudName string
	CloudinaryAPIKey    string
	CloudinaryAPISecret string
}

var AppConfig *Config

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env variables")
	}

	AppConfig = &Config{
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBName:     getEnv("DB_NAME", "hypershop"),

		Port:      getEnv("PORT", "3333"),
		GinMode:   getEnv("GIN_MODE", "debug"),
		JWTSecret: getEnv("JWT_SECRET", "default_secret"),
		AWSRegion: getEnv("AWS_REGION", "ap-south-1"),
		AWSKey:    getEnv("AWS_ACCESS_KEY", ""),
		AWSSecret: getEnv("AWS_SECRET_KEY", ""),

		CloudinaryCloudName: getEnv("CLOUDINARY_CLOUD_NAME", ""),
		CloudinaryAPIKey:    getEnv("CLOUDINARY_API_KEY", ""),
		CloudinaryAPISecret: getEnv("CLOUDINARY_API_SECRET", ""),
	}

	log.Println("✓ Environment variables loaded")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
