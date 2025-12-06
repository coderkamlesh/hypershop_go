// internal/repository/user_repo.go
package repository

import (
	"context"
	"time"

	"github.com/coderkamlesh/hypershop_go/config"
	"github.com/coderkamlesh/hypershop_go/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
    collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
    return &UserRepository{
        collection: config.DB.Collection("users"),
    }
}

// Create
func (r *UserRepository) Create(user *models.User) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    
    result, err := r.collection.InsertOne(ctx, user)
    if err != nil {
        return err
    }
    
    user.ID = result.InsertedID.(bson.ObjectID)
    return nil
}

// Read All
func (r *UserRepository) FindAll() ([]models.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := r.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var users []models.User
    if err = cursor.All(ctx, &users); err != nil {
        return nil, err
    }

    return users, nil
}

// Read One
func (r *UserRepository) FindByID(id string) (*models.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objectID, _ := bson.ObjectIDFromHex(id)
    var user models.User
    
    err := r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

// Update
func (r *UserRepository) Update(id string, user *models.User) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objectID, _ := bson.ObjectIDFromHex(id)
    user.UpdatedAt = time.Now()

    update := bson.M{
        "$set": bson.M{
            "name":       user.Name,
            "email":      user.Email,
            "phone":      user.Phone,
            "role":       user.Role,
            "updated_at": user.UpdatedAt,
        },
    }

    _, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
    return err
}

// Delete
func (r *UserRepository) Delete(id string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objectID, _ := bson.ObjectIDFromHex(id)
    _, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
    return err
}
