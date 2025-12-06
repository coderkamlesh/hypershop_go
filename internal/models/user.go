// internal/models/user.go
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
    ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
    Name      string        `bson:"name" json:"name" binding:"required"`
    Email     string        `bson:"email" json:"email" binding:"required,email"`
    Phone     string        `bson:"phone" json:"phone"`
    Role      string        `bson:"role" json:"role"` // ADMIN, SELLER, CONSUMER etc
    CreatedAt time.Time     `bson:"created_at" json:"created_at"`
    UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}
