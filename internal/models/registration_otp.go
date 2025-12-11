package models

import "time"

type RegistrationOtp struct {
	ID           string    `bson:"_id" json:"id"`
	Mobile       string    `bson:"mobile" json:"mobile"`
	OTP          string    `bson:"otp" json:"-"`
	AttemptCount int       `bson:"attempt_count" json:"attempt_count"`
	Status       bool      `bson:"status" json:"status"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
	ExpiredAt    time.Time `bson:"expired_at" json:"expired_at"`
}
