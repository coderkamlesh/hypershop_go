// repository/user_session_repository.go
package repository

import (
	"errors"
	"time"

	"github.com/coderkamlesh/hypershop_go/internal/models"
	"github.com/coderkamlesh/hypershop_go/internal/utils"
	"gorm.io/gorm"
)

type UserSessionRepository interface {
    Create(session *models.UserSession) error
    FindByToken(token string) (*models.UserSession, error)
    UpdateLastUsed(sessionID string) error
    DeactivateAllByUserID(userID string) error
}

type userSessionRepository struct {
    db *gorm.DB
}

func NewUserSessionRepository(db *gorm.DB) UserSessionRepository {
    return &userSessionRepository{db: db}
}

func (r *userSessionRepository) Create(session *models.UserSession) error {
    session.ID = utils.GenerateID("SESS")
    session.CreatedAt = time.Now()
    session.LastUsedAt = time.Now()
    session.IsActive = true
    return r.db.Create(session).Error
}

func (r *userSessionRepository) FindByToken(token string) (*models.UserSession, error) {
    var session models.UserSession
    err := r.db.Where("token = ? AND is_active = ?", token, true).First(&session).Error
    
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil // Return nil without error if not found
    }
    
    return &session, err
}

func (r *userSessionRepository) UpdateLastUsed(sessionID string) error {
    result := r.db.Model(&models.UserSession{}).
        Where("id = ?", sessionID).
        Update("last_used_at", time.Now())

    if result.Error != nil {
        return result.Error
    }
    
    if result.RowsAffected == 0 {
        return errors.New("session not found")
    }
    
    return nil
}

func (r *userSessionRepository) DeactivateAllByUserID(userID string) error {
    result := r.db.Model(&models.UserSession{}).
        Where("user_id = ?", userID).
        Update("is_active", false)

    // No error if zero rows affected - user might have no sessions
    return result.Error
}
