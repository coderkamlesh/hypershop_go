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

	Port       string
	GinMode    string
	JWTSecret  string
	AWSRegion  string
	AWSKey     string
	AWSSecret  string
	TursoToken string
	TursoUrl   string
}

var AppConfig *Config

func LoadEnv() {
	// Local development ke liye .env load karo
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env variables")
	}

	AppConfig = &Config{
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBHost:     getEnv("DB_HOST", ""),
		DBPort:     getEnv("DB_PORT", "4000"), // ✅ Default TiDB port
		DBName:     getEnv("DB_NAME", "hypershop"),

		Port:       getEnv("PORT", "8080"),
		GinMode:    getEnv("GIN_MODE", "debug"),
		JWTSecret:  getEnv("JWT_SECRET", "default_secret"),
		AWSRegion:  getEnv("AWS_REGION", "ap-south-1"),
		AWSKey:     getEnv("AWS_ACCESS_KEY", ""),
		AWSSecret:  getEnv("AWS_SECRET_KEY", ""),
		TursoToken: getEnv("TURSO_TOKEN", ""),
		TursoUrl:   getEnv("TURSO_URL", ""),
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
