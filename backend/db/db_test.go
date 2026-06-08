package db

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ilahazs/yt-webui/backend/config"
)

func TestInitAndMigrations(t *testing.T) {
	// Create a temporary directory within the test environment
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test.db")

	cfg := &config.Config{
		DBPath: dbPath,
	}

	// 1. Initial run: Database file should be created and all migrations applied
	db, err := Init(cfg)
	if err != nil {
		t.Fatalf("failed to initialize database: %v", err)
	}

	// Check if database file exists
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		t.Errorf("expected database file to exist at %s, but it does not", dbPath)
	}

	// Verify the tables exist by executing a simple count query
	tables := []string{"jobs", "job_events", "files", "settings", "schema_migrations"}
	for _, table := range tables {
		query := "SELECT count(*) FROM " + table
		var count int
		if err := db.QueryRow(query).Scan(&count); err != nil {
			t.Errorf("failed to query table %s: %v", table, err)
		}
	}

	// Verify schema_migrations table has version 1 recorded
	var version int
	err = db.QueryRow("SELECT version FROM schema_migrations WHERE version = 1").Scan(&version)
	if err != nil {
		t.Errorf("expected migration version 1 to be recorded in schema_migrations: %v", err)
	}

	// Close database connection
	if err := db.Close(); err != nil {
		t.Fatalf("failed to close database: %v", err)
	}

	// 2. Repeat run: Re-initializing should be successful and idempotent (no double migration error)
	db2, err := Init(cfg)
	if err != nil {
		t.Fatalf("failed to re-initialize database: %v", err)
	}
	defer db2.Close()

	// Verify foreign key constraint enforcement works (PRAGMA foreign_keys = ON)
	// Attempt to insert a job event referencing a non-existent job
	_, err = db2.Exec("INSERT INTO job_events (job_id, type, payload, created_at) VALUES ('non-existent-job', 'log', '{}', datetime('now'))")
	if err == nil {
		t.Error("expected foreign key constraint violation error, but query succeeded without error")
	}
}
