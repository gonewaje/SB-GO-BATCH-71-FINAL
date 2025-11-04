package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	JWTSecret   string
	DatabaseURL string
}

func Load() *Config {
	// Load local .env but allow host env to override
	_ = godotenv.Load("config/.env")

	if _, err := os.Stat("config/.env"); err == nil {
		log.Println("✅ Loaded config/.env file (host env can override)")
	} else {
		log.Println("📦 No local .env found — using system environment variables")
	}

	cfg := &Config{
		Port:       getenv("PORT", "8080"),
		DBHost:     getenv("DB_HOST", "localhost"),
		DBPort:     getenv("DB_PORT", "5432"),
		DBUser:     getenv("DB_USER", "postgres"),
		DBPassword: getenv("DB_PASSWORD", ""),
		DBName:     getenv("DB_NAME", "postgres"),
		JWTSecret:  getenv("JWT_SECRET", "supersecret"),
	}

	cfg.DatabaseURL = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	log.Printf("🔧 Loaded configuration: DB_HOST=%s DB_NAME=%s PORT=%s",
		cfg.DBHost, cfg.DBName, cfg.Port)

	return cfg
}

func getenv(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
