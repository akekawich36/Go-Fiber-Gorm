package configs

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnect() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal("Error loading config")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=%s TimeZone=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.DBName,
		config.Database.SSLMode,
		config.Database.Timezone)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get DB instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	log.Println("✅ Connected to PostgreSQL")
}
