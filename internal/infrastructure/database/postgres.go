package database

import (
	"fmt"
	"log"
	"time"

	"github.com/akekawich36/go-authen/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnect() (*gorm.DB, error) {
	config, err := configs.LoadConfig()
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
	return DB, nil
}
