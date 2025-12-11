package indexes

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// CreateOTPIndexes creates indexes for OTP collections
func CreateOTPIndexes(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ========== REGISTRATION_OTPS COLLECTION ==========
	regOtpCollection := db.Collection("registration_otps")

	regOtpIndexes := []mongo.IndexModel{
		// Unique index on mobile
		{
			Keys:    bson.D{{Key: "mobile", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		// TTL index - auto-delete after 10 minutes
		{
			Keys:    bson.D{{Key: "expired_at", Value: 1}},
			Options: options.Index().SetExpireAfterSeconds(0),
		},
	}

	_, err := regOtpCollection.Indexes().CreateMany(ctx, regOtpIndexes)
	if err != nil {
		log.Printf("❌ Error creating registration OTP indexes: %v", err)
		return err
	}

	// ========== USER_OTPS COLLECTION ==========
	userOtpCollection := db.Collection("user_otps")

	userOtpIndexes := []mongo.IndexModel{
		// Unique index on mobile
		{
			Keys:    bson.D{{Key: "mobile", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		// Index on user_id
		{
			Keys: bson.D{{Key: "user_id", Value: 1}},
		},
		// TTL index - auto-delete after 10 minutes
		{
			Keys:    bson.D{{Key: "expired_at", Value: 1}},
			Options: options.Index().SetExpireAfterSeconds(0),
		},
	}

	_, err = userOtpCollection.Indexes().CreateMany(ctx, userOtpIndexes)
	if err != nil {
		log.Printf("❌ Error creating user OTP indexes: %v", err)
		return err
	}

	log.Println("✅ OTP indexes created")
	return nil
}
