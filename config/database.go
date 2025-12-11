package config

import (
	"fmt"
	"log"
	"time"

	"github.com/coderkamlesh/hypershop_go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
    // CockroachDB connection string
    dsn := fmt.Sprintf(
        "postgresql://%s:%s@%s/%s?sslmode=require",
        AppConfig.DBUser,
        AppConfig.DBPassword,
        AppConfig.DBHost,
        AppConfig.DBName,
    )

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
        NowFunc: func() time.Time {
            return time.Now().UTC()
        },
    })

    if err != nil {
        log.Fatal("❌ Postgres connection failed:", err)
    }

    // Connection pool settings
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatal("❌ Failed to configure connection pool:", err)
    }

    sqlDB.SetMaxOpenConns(25)
    sqlDB.SetMaxIdleConns(5)
    sqlDB.SetConnMaxLifetime(5 * time.Minute)

    fmt.Println("✓ CockroachDB Connected!")

    // Auto migrate tables
    if err := AutoMigrate(); err != nil {
        log.Printf("⚠️ Warning: Migration error: %v", err)
    }
}

func AutoMigrate() error {
    return DB.AutoMigrate(
        &models.User{},
        &models.UserSession{},
        &models.UserOtp{},
        &models.RegistrationOtp{},
    )
}
