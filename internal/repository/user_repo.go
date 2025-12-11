package repository

import (
	"context"
	"errors"
	"time"

	"github.com/coderkamlesh/hypershop_go/internal/models"
	"github.com/coderkamlesh/hypershop_go/internal/utils"
	"go.mongodb.org/mongo-driver/v2/bson"          // ✅ v2
	"go.mongodb.org/mongo-driver/v2/mongo"         // ✅ v2
	"go.mongodb.org/mongo-driver/v2/mongo/options" // ✅ v2
)

type UserRepository interface {
	Create(user *models.User) error
	FindByID(id string) (*models.User, error)
	FindByMobile(mobile string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindAll(page, pageSize int) ([]*models.User, int64, error)
	FindByRole(role string, page, pageSize int) ([]*models.User, int64, error)
	Update(id string, user *models.User) error
	Delete(id string) error
	ExistsByMobile(mobile string) (bool, error)
	ExistsByEmail(email string) (bool, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("users"),
	}
}

// Create creates a new user
func (r *userRepository) Create(user *models.User) error {
	// Generate custom ID
	user.ID = utils.GenerateUserID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsActive = true

	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

// FindByID finds user by ID
func (r *userRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByMobile finds user by mobile number
func (r *userRepository) FindByMobile(mobile string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.Background(), bson.M{"mobile": mobile}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, nil // Return nil without error if not found
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByEmail finds user by email
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindAll gets all users with pagination
func (r *userRepository) FindAll(page, pageSize int) ([]*models.User, int64, error) {
	skip := (page - 1) * pageSize

	// Count total documents
	total, err := r.collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return nil, 0, err
	}

	// Find with pagination
	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(pageSize)).
		SetSort(bson.D{{Key: "created_at", Value: -1}}) // Latest first

	cursor, err := r.collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())

	var users []*models.User
	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// FindByRole finds users by role with pagination
func (r *userRepository) FindByRole(role string, page, pageSize int) ([]*models.User, int64, error) {
	skip := (page - 1) * pageSize
	filter := bson.M{"role": role}

	// Count total
	total, err := r.collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, 0, err
	}

	// Find with pagination
	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(pageSize)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())

	var users []*models.User
	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Update updates user details
func (r *userRepository) Update(id string, user *models.User) error {
	user.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"role":       user.Role,
			"is_active":  user.IsActive,
			"updated_at": user.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		update,
	)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// Delete deletes a user (soft delete by setting is_active = false)
func (r *userRepository) Delete(id string) error {
	update := bson.M{
		"$set": bson.M{
			"is_active":  false,
			"updated_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		update,
	)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// ExistsByMobile checks if user exists with given mobile
func (r *userRepository) ExistsByMobile(mobile string) (bool, error) {
	count, err := r.collection.CountDocuments(context.Background(), bson.M{"mobile": mobile})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ExistsByEmail checks if user exists with given email
func (r *userRepository) ExistsByEmail(email string) (bool, error) {
	count, err := r.collection.CountDocuments(context.Background(), bson.M{"email": email})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
