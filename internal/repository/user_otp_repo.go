package repository

import (
	"context"
	"time"

	"github.com/coderkamlesh/hypershop_go/internal/models"
	"github.com/coderkamlesh/hypershop_go/internal/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserOtpRepository interface {
	Create(otp *models.UserOtp) error
	FindByMobile(mobile string) (*models.UserOtp, error)
	Update(otp *models.UserOtp) error
}

type userOtpRepository struct {
	collection *mongo.Collection
}

func NewUserOtpRepository(db *mongo.Database) UserOtpRepository {
	return &userOtpRepository{
		collection: db.Collection("user_otps"),
	}
}

func (r *userOtpRepository) Create(otp *models.UserOtp) error {
	otp.ID = utils.GenerateID("UOTP")
	otp.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(context.Background(), otp)
	return err
}

func (r *userOtpRepository) FindByMobile(mobile string) (*models.UserOtp, error) {
	var otp models.UserOtp
	err := r.collection.FindOne(context.Background(), bson.M{"mobile": mobile}).Decode(&otp)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &otp, err
}

func (r *userOtpRepository) Update(otp *models.UserOtp) error {
	update := bson.M{
		"$set": bson.M{
			"otp":           otp.OTP,
			"user_id":       otp.UserID,
			"attempt_count": otp.AttemptCount,
			"status":        otp.Status,
			"created_at":    otp.CreatedAt,
			"expired_at":    otp.ExpiredAt,
		},
	}
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": otp.ID}, update)
	return err
}
