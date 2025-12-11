package indexes

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// CreateSessionIndexes creates indexes for user_sessions collection
func CreateSessionIndexes(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection("user_sessions")

	indexes := []mongo.IndexModel{
		// Index on user_id
		{
			Keys: bson.D{{Key: "user_id", Value: 1}},
		},
		// Index on token (for quick lookup)
		{
			Keys: bson.D{{Key: "token", Value: 1}},
		},
		// Compound index on user_id + is_active
		{
			Keys: bson.D{
				{Key: "user_id", Value: 1},
				{Key: "is_active", Value: 1},
			},
		},
		// Index on created_at
		{
			Keys: bson.D{{Key: "created_at", Value: -1}},
		},
	}

	_, err := collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		log.Printf("❌ Error creating session indexes: %v", err)
		return err
	}

	log.Println("✅ Session indexes created")
	return nil
}
