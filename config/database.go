package config

import (
	"database/sql" // Standard sql package zaroori hai
	"fmt"
	"log"
	"time"

	_ "github.com/tursodatabase/libsql-client-go/libsql" // Turso Driver
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	start := time.Now()

	// 1. Turso DSN create karo
	// Make sure AppConfig.TursoUrl me "libsql://" prefix ho
	dsn := fmt.Sprintf("%s?authToken=%s", AppConfig.TursoUrl, AppConfig.TursoToken)

	// 2. Pehle standard SQL connection open karo
	sqlDB, err := sql.Open("libsql", dsn)
	if err != nil {
		log.Fatal("Failed to open connection to Turso:", err)
	}

	// 3. Ab GORM ko us connection ke saath initialize karo
	DB, err = gorm.Open(sqlite.New(sqlite.Config{
		Conn: sqlDB, // Yaha hum apna Turso connection pass kar rahe hain
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		log.Fatal("Turso connection via GORM failed:", err)
	}

	// Connection pool settings (Turso HTTP based hai par pool maintain karna acha hai)
	// Note: Hum 'sqlDB' variable already upar bana chuke hain, dubara DB.DB() ki zarurat nahi

	// Turso generally high concurrency handle kar leta hai,
	// lekin agar serverless/lambda environment hai to connections kam rakho.
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	fmt.Println("✓ Connected to Turso DB successfully!")
	duration := time.Since(start)
	fmt.Printf("⏱️ DB Connection took: %v\n", duration)
}
