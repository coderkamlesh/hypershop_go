package indexes

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// CreateAllIndexes creates indexes for all collections
func CreateAllIndexes(db *mongo.Database) error {
	// User indexes
	if err := CreateUserIndexes(db); err != nil {
		return err
	}

	// OTP indexes
	if err := CreateOTPIndexes(db); err != nil {
		return err
	}

	// Session indexes
	if err := CreateSessionIndexes(db); err != nil {
		return err
	}

	// Future: Product, Order indexes
	// if err := CreateProductIndexes(db); err != nil {
	//     return err
	// }

	return nil
}
