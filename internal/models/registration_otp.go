package models

import "time"

type RegistrationOtp struct {
    ID           string    `gorm:"primaryKey;type:varchar(50)" json:"id"`
    Mobile       string    `gorm:"type:varchar(15);uniqueIndex;not null" json:"mobile"`
    OTP          string    `gorm:"type:varchar(6);not null" json:"-"`
    AttemptCount int       `gorm:"default:0" json:"attempt_count"`
    Status       bool      `gorm:"default:false" json:"status"`
    CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
    ExpiredAt    time.Time `gorm:"index" json:"expired_at"`
}