package repository

import (
	"context"
	"time"

	"github.com/coderkamlesh/hypershop_go/internal/models"
	"github.com/coderkamlesh/hypershop_go/internal/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserSessionRepository interface {
	Create(session *models.UserSession) error
	FindByToken(token string) (*models.UserSession, error)
	UpdateLastUsed(sessionID string) error
	DeactivateAllByUserID(userID string) error
}

type userSessionRepository struct {
	collection *mongo.Collection
}

func NewUserSessionRepository(db *mongo.Database) UserSessionRepository {
	return &userSessionRepository{
		collection: db.Collection("user_sessions"),
	}
}

func (r *userSessionRepository) Create(session *models.UserSession) error {
	session.ID = utils.GenerateID("SESS")
	session.CreatedAt = time.Now()
	session.LastUsedAt = time.Now()
	session.IsActive = true
	_, err := r.collection.InsertOne(context.Background(), session)
	return err
}

func (r *userSessionRepository) FindByToken(token string) (*models.UserSession, error) {
	var session models.UserSession
	err := r.collection.FindOne(context.Background(), bson.M{"token": token, "is_active": true}).Decode(&session)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &session, err
}

func (r *userSessionRepository) UpdateLastUsed(sessionID string) error {
	update := bson.M{
		"$set": bson.M{
			"last_used_at": time.Now(),
		},
	}
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": sessionID}, update)
	return err
}

func (r *userSessionRepository) DeactivateAllByUserID(userID string) error {
	update := bson.M{
		"$set": bson.M{
			"is_active": false,
		},
	}
	_, err := r.collection.UpdateMany(context.Background(), bson.M{"user_id": userID}, update)
	return err
}
