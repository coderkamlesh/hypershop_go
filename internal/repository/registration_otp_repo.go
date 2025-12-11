// repository/registration_otp_repository.go
package repository

import (
	"errors"
	"time"

	"github.com/coderkamlesh/hypershop_go/internal/models"
	"github.com/coderkamlesh/hypershop_go/internal/utils"
	"gorm.io/gorm"
)

type RegistrationOtpRepository interface {
    Create(otp *models.RegistrationOtp) error
    FindByMobile(mobile string) (*models.RegistrationOtp, error)
    Update(otp *models.RegistrationOtp) error
}

type registrationOtpRepository struct {
    db *gorm.DB
}

func NewRegistrationOtpRepository(db *gorm.DB) RegistrationOtpRepository {
    return &registrationOtpRepository{db: db}
}

func (r *registrationOtpRepository) Create(otp *models.RegistrationOtp) error {
    otp.ID = utils.GenerateID("ROTP")
    otp.CreatedAt = time.Now()
    return r.db.Create(otp).Error
}

func (r *registrationOtpRepository) FindByMobile(mobile string) (*models.RegistrationOtp, error) {
    var otp models.RegistrationOtp
    err := r.db.Where("mobile = ?", mobile).First(&otp).Error
    
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil // Return nil without error if not found
    }
    
    return &otp, err
}

func (r *registrationOtpRepository) Update(otp *models.RegistrationOtp) error {
    result := r.db.Model(&models.RegistrationOtp{}).
        Where("id = ?", otp.ID).
        Updates(map[string]interface{}{
            "otp":           otp.OTP,
            "attempt_count": otp.AttemptCount,
            "status":        otp.Status,
            "created_at":    otp.CreatedAt,
            "expired_at":    otp.ExpiredAt,
        })

    if result.Error != nil {
        return result.Error
    }
    
    if result.RowsAffected == 0 {
        return errors.New("registration OTP not found")
    }
    
    return nil
}
