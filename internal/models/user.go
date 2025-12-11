package models

import (
	"time"

	"github.com/coderkamlesh/hypershop_go/internal/constants"
)

type User struct {
    ID        string         `gorm:"primaryKey;type:varchar(50)" json:"id"`
    Name      string         `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
    Mobile    string         `gorm:"type:varchar(15);uniqueIndex;not null" json:"mobile" binding:"required"`
    Email     string         `gorm:"type:varchar(100);uniqueIndex" json:"email"`
    Role      constants.Role `gorm:"type:varchar(20);default:'customer'" json:"role"`
    Password  string         `gorm:"type:varchar(255)" json:"-"`
    IsActive  bool           `gorm:"default:true" json:"is_active"`
    CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}