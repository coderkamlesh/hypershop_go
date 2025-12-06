// config/env.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    MongoURI   string
    DBName     string
    Port       string
    GinMode    string
    JWTSecret  string
    AWSRegion  string
    AWSKey     string
    AWSSecret  string
}

var AppConfig *Config

func LoadEnv() {
    // .env load karo
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  No .env file found, using system env variables")
    }

    AppConfig = &Config{
        MongoURI:   getEnv("MONGODB_URI", ""),
        DBName:     getEnv("DB_NAME", "hypershop"),
        Port:       getEnv("PORT", "8080"),
        GinMode:    getEnv("GIN_MODE", "debug"),
        JWTSecret:  getEnv("JWT_SECRET", "default_secret"),
        AWSRegion:  getEnv("AWS_REGION", "ap-south-1"),
        AWSKey:     getEnv("AWS_ACCESS_KEY", ""),
        AWSSecret:  getEnv("AWS_SECRET_KEY", ""),
    }

    log.Println("✓ Environment variables loaded")
}

// Helper: default value support
func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}
