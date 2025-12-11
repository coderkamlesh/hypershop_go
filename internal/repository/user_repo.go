// repository/user_repository.go
package repository

import (
	"errors"
	"time"

	"github.com/coderkamlesh/hypershop_go/internal/models"
	"github.com/coderkamlesh/hypershop_go/internal/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
    Create(user *models.User) error
    FindByID(id string) (*models.User, error)
    FindByMobile(mobile string) (*models.User, error)
    FindByEmail(email string) (*models.User, error)
    FindAll(page, pageSize int) ([]*models.User, int64, error)
    Update(id string, user *models.User) error
    Delete(id string) error
    ExistsByMobile(mobile string) (bool, error)
    ExistsByEmail(email string) (bool, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
    user.ID = utils.GenerateUserID()
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    user.IsActive = true
    return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id string) (*models.User, error) {
    var user models.User
    err := r.db.Where("id = ?", id).First(&user).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, errors.New("user not found")
    }
    return &user, err
}

func (r *userRepository) FindByMobile(mobile string) (*models.User, error) {
    var user models.User
    err := r.db.Where("mobile = ?", mobile).First(&user).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    return &user, err
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.db.Where("email = ?", email).First(&user).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    return &user, err
}

func (r *userRepository) FindAll(page, pageSize int) ([]*models.User, int64, error) {
    var users []*models.User
    var total int64

    offset := (page - 1) * pageSize

    if err := r.db.Model(&models.User{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }

    err := r.db.Offset(offset).Limit(pageSize).
        Order("created_at DESC").
        Find(&users).Error

    return users, total, err
}

func (r *userRepository) Update(id string, user *models.User) error {
    user.UpdatedAt = time.Now()
    
    result := r.db.Model(&models.User{}).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "name":       user.Name,
            "email":      user.Email,
            "role":       user.Role,
            "is_active":  user.IsActive,
            "updated_at": user.UpdatedAt,
        })

    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("user not found")
    }
    return nil
}

func (r *userRepository) Delete(id string) error {
    result := r.db.Model(&models.User{}).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "is_active":  false,
            "updated_at": time.Now(),
        })

    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("user not found")
    }
    return nil
}

func (r *userRepository) ExistsByMobile(mobile string) (bool, error) {
    var count int64
    err := r.db.Model(&models.User{}).Where("mobile = ?", mobile).Count(&count).Error
    return count > 0, err
}

func (r *userRepository) ExistsByEmail(email string) (bool, error) {
    var count int64
    err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
    return count > 0, err
}
