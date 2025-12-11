package indexes

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// CreateUserIndexes creates indexes for users collection
func CreateUserIndexes(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection("users")

	indexes := []mongo.IndexModel{
		// Unique index on mobile
		{
			Keys:    bson.D{{Key: "mobile", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		// Index on role
		{
			Keys: bson.D{{Key: "role", Value: 1}},
		},
		// Index on created_at
		{
			Keys: bson.D{{Key: "created_at", Value: -1}},
		},
		// Unique sparse index on email
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true).SetSparse(true),
		},
		// Compound index on role + created_at
		{
			Keys: bson.D{
				{Key: "role", Value: 1},
				{Key: "created_at", Value: -1},
			},
		},
	}

	_, err := collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		log.Printf("❌ Error creating user indexes: %v", err)
		return err
	}

	log.Println("✅ User indexes created")
	return nil
}
