package models

import (
	"time"

	"github.com/coderkamlesh/hypershop_go/internal/constants"
)

type User struct {
	ID        string         `bson:"_id" json:"id"`
	Name      string         `bson:"name" json:"name" binding:"required"`
	Mobile    string         `bson:"mobile" json:"mobile" binding:"required"`
	Email     string         `bson:"email" json:"email"`
	Role      constants.Role `bson:"role" json:"role"`
	Password  string         `bson:"password" json:"-"`
	IsActive  bool           `bson:"is_active" json:"is_active"`
	CreatedAt time.Time      `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time      `bson:"updated_at" json:"updated_at"`
}
