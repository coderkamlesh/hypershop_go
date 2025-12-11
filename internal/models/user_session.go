package models

import (
	"time"
)

type UserSession struct {
    ID         string    `gorm:"primaryKey;type:varchar(50)" json:"id"`
    UserID     string    `gorm:"type:varchar(50);index;not null" json:"user_id"`
    Token      string    `gorm:"type:text;uniqueIndex;not null" json:"-"`
    DeviceInfo string    `gorm:"type:varchar(255)" json:"device_info"`
    Source     string    `gorm:"type:varchar(20)" json:"source"`
    IsActive   bool      `gorm:"default:true;index" json:"is_active"`
    CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
    LastUsedAt time.Time `json:"last_used_at"`
}