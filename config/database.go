// config/database.go
package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // .env se URI load
    clientOptions := options.Client().ApplyURI(AppConfig.MongoURI)
    client, err := mongo.Connect(clientOptions)
    if err != nil {
        log.Fatal("❌ MongoDB connection failed:", err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("❌ MongoDB ping failed:", err)
    }

    DB = client.Database(AppConfig.DBName)
    fmt.Printf("✓ MongoDB Connected to '%s' database!\n", AppConfig.DBName)
}
