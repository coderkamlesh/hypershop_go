package repository

import (
	"context"
	"time"

	"github.com/coderkamlesh/hypershop_go/internal/models"
	"github.com/coderkamlesh/hypershop_go/internal/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RegistrationOtpRepository interface {
	Create(otp *models.RegistrationOtp) error
	FindByMobile(mobile string) (*models.RegistrationOtp, error)
	Update(otp *models.RegistrationOtp) error
}

type registrationOtpRepository struct {
	collection *mongo.Collection
}

func NewRegistrationOtpRepository(db *mongo.Database) RegistrationOtpRepository {
	return &registrationOtpRepository{
		collection: db.Collection("registration_otps"),
	}
}

func (r *registrationOtpRepository) Create(otp *models.RegistrationOtp) error {
	otp.ID = utils.GenerateID("ROTP")
	otp.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(context.Background(), otp)
	return err
}

func (r *registrationOtpRepository) FindByMobile(mobile string) (*models.RegistrationOtp, error) {
	var otp models.RegistrationOtp
	err := r.collection.FindOne(context.Background(), bson.M{"mobile": mobile}).Decode(&otp)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &otp, err
}

func (r *registrationOtpRepository) Update(otp *models.RegistrationOtp) error {
	update := bson.M{
		"$set": bson.M{
			"otp":           otp.OTP,
			"attempt_count": otp.AttemptCount,
			"status":        otp.Status,
			"created_at":    otp.CreatedAt,
			"expired_at":    otp.ExpiredAt,
		},
	}
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": otp.ID}, update)
	return err
}
