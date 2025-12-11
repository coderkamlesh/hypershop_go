package indexes

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// CreateProductIndexes creates indexes for products collection
func CreateProductIndexes(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.Collection("products")

	indexes := []mongo.IndexModel{
		// Text index for search
		{
			Keys: bson.D{
				{Key: "name", Value: "text"},
				{Key: "description", Value: "text"},
			},
		},
		// Index on category_id
		{
			Keys: bson.D{{Key: "category_id", Value: 1}},
		},
		// Compound index on category + price
		{
			Keys: bson.D{
				{Key: "category_id", Value: 1},
				{Key: "price", Value: 1},
			},
		},
		// Index on seller_id
		{
			Keys: bson.D{{Key: "seller_id", Value: 1}},
		},
		// Index on is_active (to filter active products)
		{
			Keys: bson.D{{Key: "is_active", Value: 1}},
		},
	}

	_, err := collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		log.Printf("❌ Error creating product indexes: %v", err)
		return err
	}

	log.Println("✅ Product indexes created")
	return nil
}
