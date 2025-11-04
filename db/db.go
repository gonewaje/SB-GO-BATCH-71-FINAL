package db

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Open(dsn string) *sql.DB {
	log.Println("🔌 Connecting to database...")
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to open DB: %v", err)
	}
	if err := database.Ping(); err != nil {
		log.Fatalf("❌ Failed to ping DB: %v", err)
	}
	log.Println("✅ Connected to database")

	if err := runMigration(database); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	return database
}

func runMigration(database *sql.DB) error {
	mig := filepath.Join("db", "sql_migrations", "migrate.sql")
	content, err := ioutil.ReadFile(mig)
	if err != nil {
		log.Printf("⚠️ Could not read migration file at %s; skipping", mig)
		return nil
	}
	up := extractUp(string(content))
	if up == "" {
		log.Println("⚠️ No '-- +migrate Up' section found; skipping")
		return nil
	}
	log.Println("🚀 Running migration...")
	if _, err := database.Exec(up); err != nil {
		// If exist errors happen, ignore them; this is idempotent-ish migration
		if !strings.Contains(strings.ToLower(err.Error()), "already exists") {
			return err
		}
	}
	log.Println("✅ Migration complete")
	return nil
}

func extractUp(s string) string {
	parts := strings.Split(s, "-- +migrate Up")
	if len(parts) < 2 {
		return ""
	}
	up := strings.Split(parts[1], "-- +migrate Down")
	return strings.TrimSpace(up[0])
}
