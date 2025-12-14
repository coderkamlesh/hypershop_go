package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	start := time.Now()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=skip-verify",
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		log.Fatal("MySQL connection failed:", err)
	}

	// Connection pool settings
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to configure connection pool:", err)
	}

	sqlDB.SetMaxOpenConns(10)

	sqlDB.SetMaxIdleConns(5)

	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("✓ Connected to MySQL DB successfully!")
	duration := time.Since(start)
	fmt.Printf("⏱️ DB Connection took: %v\n", duration)
}
