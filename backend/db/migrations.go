package db

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

var migrationFileRegex = regexp.MustCompile(`^(\d{4})_.*\.sql$`)

type migration struct {
	version  int
	filename string
	content  string
}

// RunMigrations checks the applied migrations in the schema_migrations table,
// reads the embedded migration SQL files, sorts them, and applies pending migrations.
func RunMigrations(db *sql.DB) error {
	// 1. Ensure schema_migrations exists
	createSchemaTableSQL := `
	CREATE TABLE IF NOT EXISTS schema_migrations (
		version INTEGER PRIMARY KEY,
		applied_at DATETIME NOT NULL
	);`
	if _, err := db.Exec(createSchemaTableSQL); err != nil {
		return fmt.Errorf("failed to ensure schema_migrations table exists: %w", err)
	}

	// 2. Fetch applied migrations
	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return fmt.Errorf("failed to query applied migrations: %w", err)
	}
	defer rows.Close()

	applied := make(map[int]bool)
	for rows.Next() {
		var version int
		if err := rows.Scan(&version); err != nil {
			return fmt.Errorf("failed to scan applied migration version: %w", err)
		}
		applied[version] = true
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("error reading migration versions: %w", err)
	}

	// 3. Read and parse embedded migration files
	entries, err := migrationsFS.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("failed to read embedded migrations directory: %w", err)
	}

	var pending []migration
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		matches := migrationFileRegex.FindStringSubmatch(filename)
		if len(matches) != 2 {
			// Skip files that do not match the expected naming pattern (e.g. README or temp files)
			continue
		}

		version, err := strconv.Atoi(matches[1])
		if err != nil {
			return fmt.Errorf("failed to parse version from migration filename %s: %w", filename, err)
		}

		if applied[version] {
			continue
		}

		content, err := migrationsFS.ReadFile("migrations/" + filename)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", filename, err)
		}

		pending = append(pending, migration{
			version:  version,
			filename: filename,
			content:  string(content),
		})
	}

	// 4. Sort pending migrations by version ascending
	sort.Slice(pending, func(i, j int) bool {
		return pending[i].version < pending[j].version
	})

	// 5. Execute migrations in order
	for _, m := range pending {
		log.Printf("Applying database migration: %s (version %d)...", m.filename, m.version)
		if err := applyMigration(db, m); err != nil {
			return fmt.Errorf("failed to apply migration %s: %w", m.filename, err)
		}
		log.Printf("Successfully applied database migration: %s", m.filename)
	}

	return nil
}

func applyMigration(db *sql.DB, m migration) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Execute migration SQL
	if _, err := tx.Exec(m.content); err != nil {
		return fmt.Errorf("failed to execute SQL statements: %w", err)
	}

	// Record applied version
	insertSQL := `INSERT INTO schema_migrations (version, applied_at) VALUES (?, datetime('now'))`
	if _, err := tx.Exec(insertSQL, m.version); err != nil {
		return fmt.Errorf("failed to record applied version: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
