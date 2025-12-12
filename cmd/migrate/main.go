package main

import (
	"fmt"
	"log"

	"github.com/coderkamlesh/hypershop_go/config"
	"github.com/coderkamlesh/hypershop_go/internal/models"
)

func main() {
	fmt.Println("🚀 Running migrations...")
	config.LoadEnv()
	// Connect DB
	config.ConnectDB()

	// Run migrations
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.UserOtp{},
		&models.RegistrationOtp{},
		&models.UserSession{},
		// future tables add here: &models.Product{}, &models.Order{}, etc.
	)

	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}

	fmt.Println("✅ Migration completed successfully!")
}
