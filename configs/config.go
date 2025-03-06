package configs

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type AppConfig struct {
	Name string
	Env  string
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	Timezone string
}

type JWTConfig struct {
	Secret             string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessTokenExpiry, err := time.ParseDuration(getEnv("JWT_TOKEN_EXPIRE", "12h"))
	if err != nil {
		return nil, err
	}

	refreshTokenExpiry, err := time.ParseDuration(getEnv("JWT_REFRESH_TOKEN_EXPIRE", "15m"))
	if err != nil {
		return nil, err
	}

	return &Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "auth-api"),
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "auth_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
			Timezone: getEnv("DB_TIMEZONE", "UTC"),
		},
		JWT: JWTConfig{
			Secret:             getEnv("JWT_SECRET", "your-secret-key"),
			AccessTokenExpiry:  accessTokenExpiry,
			RefreshTokenExpiry: refreshTokenExpiry,
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
