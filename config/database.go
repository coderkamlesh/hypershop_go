package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	// connection string
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=require",
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBHost,
		AppConfig.DBName,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
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

	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(30 * time.Second)
	sqlDB.SetConnMaxIdleTime(10 * time.Second)

	fmt.Println("✓ DB Connected!")

}
