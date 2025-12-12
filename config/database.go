// package config

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// var DB *gorm.DB

// func ConnectDB() {
// 	// connection string
// 	dsn := fmt.Sprintf(
// 		"postgresql://%s:%s@%s/%s?sslmode=require",
// 		AppConfig.DBUser,
// 		AppConfig.DBPassword,
// 		AppConfig.DBHost,
// 		AppConfig.DBName,
// 	)

// 	var err error
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent),
// 		NowFunc: func() time.Time {
// 			return time.Now().UTC()
// 		},
// 	})

// 	if err != nil {
// 		log.Fatal("❌ Postgres connection failed:", err)
// 	}

// 	// Connection pool settings
// 	sqlDB, err := DB.DB()
// 	if err != nil {
// 		log.Fatal("❌ Failed to configure connection pool:", err)
// 	}

// 	sqlDB.SetMaxOpenConns(1)
// 	sqlDB.SetMaxIdleConns(1)
// 	sqlDB.SetConnMaxLifetime(30 * time.Second)
// 	sqlDB.SetConnMaxIdleTime(10 * time.Second)

// 	fmt.Println("✓ DB Connected!")

// }
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
	// ✅ TiDB / MySQL Connection String (DSN)
	// Format: user:pass@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local&tls=true
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=true",
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBHost,
		AppConfig.DBPort, // Added port
		AppConfig.DBName,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Info mode rakho shuru me debug ke liye
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		log.Fatal("❌ TiDB/MySQL connection failed:", err)
	}

	// Connection pool settings
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("❌ Failed to configure connection pool:", err)
	}

	// ✅ TiDB Serverless connection handling
	// Serverless db connections ko jaldi close kar deta hai agar idle ho,
	// isliye lifetime thoda kam rakha hai.
	// Sirf 1 connection open rakho (Lambda concurrency sambhal lega)
sqlDB.SetMaxOpenConns(1) 

// Idle connection mat rakho (taaki TiDB confuse na ho)
sqlDB.SetMaxIdleConns(1)

// Connection ko har 2 minute me recycle karo
sqlDB.SetConnMaxLifetime(2 * time.Minute)

	fmt.Println("✓ Connected to TiDB (MySQL) successfully!")
	duration := time.Since(start)
    fmt.Printf("⏱️ DB Connection took: %v\n", duration)
}