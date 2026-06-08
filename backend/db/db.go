package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"

	"github.com/ilahazs/yt-webui/backend/config"
)

// Init initializes the database connection, ensures directories exist,
// configures SQLite connection settings, and applies pending migrations.
func Init(cfg *config.Config) (*sql.DB, error) {
	if cfg.DBPath == "" {
		return nil, fmt.Errorf("database path is not configured")
	}

	// 1. Ensure the parent directory of the database file exists
	dbDir := filepath.Dir(cfg.DBPath)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory %s: %w", dbDir, err)
	}

	// 2. Open SQLite database connection
	db, err := sql.Open("sqlite", cfg.DBPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// 3. Configure SQLite connection pool limits
	// SQLite performs best with limited connection pools to prevent write contention.
	// Since we enable WAL mode, multiple readers are fine, but writes are still serialized.
	db.SetMaxOpenConns(10)

	// 4. Apply SQLite-specific optimization pragmas
	pragmas := []string{
		"PRAGMA foreign_keys = ON;",
		"PRAGMA journal_mode = WAL;",
		"PRAGMA synchronous = NORMAL;",
		"PRAGMA busy_timeout = 5000;",
	}
	for _, pragma := range pragmas {
		if _, err := db.Exec(pragma); err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to execute SQLite pragma %q: %w", pragma, err)
		}
	}

	// 5. Run embedded database migrations
	if err := RunMigrations(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to run database migrations: %w", err)
	}

	return db, nil
}
