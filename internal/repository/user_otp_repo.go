// repository/user_otp_repository.go
package repository

import (
	"errors"
	"time"

	"github.com/coderkamlesh/hypershop_go/internal/models"
	"github.com/coderkamlesh/hypershop_go/internal/utils"
	"gorm.io/gorm"
)

type UserOtpRepository interface {
    Create(otp *models.UserOtp) error
    FindByMobile(mobile string) (*models.UserOtp, error)
    Update(otp *models.UserOtp) error
}

type userOtpRepository struct {
    db *gorm.DB
}

func NewUserOtpRepository(db *gorm.DB) UserOtpRepository {
    return &userOtpRepository{db: db}
}

func (r *userOtpRepository) Create(otp *models.UserOtp) error {
    otp.ID = utils.GenerateID("UOTP")
    otp.CreatedAt = time.Now()
    return r.db.Create(otp).Error
}

func (r *userOtpRepository) FindByMobile(mobile string) (*models.UserOtp, error) {
    var otp models.UserOtp
    err := r.db.Where("mobile = ?", mobile).First(&otp).Error
    
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil // Return nil without error if not found
    }
    
    return &otp, err
}

func (r *userOtpRepository) Update(otp *models.UserOtp) error {
    result := r.db.Model(&models.UserOtp{}).
        Where("id = ?", otp.ID).
        Updates(map[string]interface{}{
            "otp":           otp.OTP,
            "user_id":       otp.UserID,
            "attempt_count": otp.AttemptCount,
            "status":        otp.Status,
            "created_at":    otp.CreatedAt,
            "expired_at":    otp.ExpiredAt,
        })

    if result.Error != nil {
        return result.Error
    }
    
    if result.RowsAffected == 0 {
        return errors.New("user OTP not found")
    }
    
    return nil
}
