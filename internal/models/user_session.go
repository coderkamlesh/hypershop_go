package models

import "time"

type UserSession struct {
	ID         string    `bson:"_id" json:"id"`
	UserID     string    `bson:"user_id" json:"user_id"`
	Token      string    `bson:"token" json:"-"`
	DeviceInfo string    `bson:"device_info" json:"device_info"`
	Source     string    `bson:"source" json:"source"` // "android", "ios", "web"
	IsActive   bool      `bson:"is_active" json:"is_active"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
	LastUsedAt time.Time `bson:"last_used_at" json:"last_used_at"`
}
