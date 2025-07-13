// Copyright 2012-2025 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package migration handles database migrations for the NATS server
package migration

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-gorp/gorp/v3"
)

// Migration represents a database migration
type Migration struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	AppliedAt time.Time `db:"applied_at"`
}

// DBExecutor represents a database executor that can execute migrations
type DBExecutor interface {
	Execute(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// Migrator handles the migration process
type Migrator struct {
	db     *gorp.DbMap
	dbType string
}

// NewMigrator creates a new migrator
func NewMigrator(db *gorp.DbMap, dbType string) *Migrator {
	return &Migrator{
		db:     db,
		dbType: dbType,
	}
}

// Apply applies the specified migrations
func (m *Migrator) Apply(migrations []string) error {
	if m.db == nil {
		return errors.New("database connection is nil")
	}

	// Create migrations table if it doesn't exist
	if err := m.createMigrationsTable(); err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Get applied migrations
	appliedMigrations, err := m.getAppliedMigrations()
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	// Apply new migrations
	for _, migration := range migrations {
		if !m.isMigrationApplied(appliedMigrations, migration) {
			if err := m.applyMigration(migration); err != nil {
				return fmt.Errorf("failed to apply migration %s: %w", migration, err)
			}
		}
	}

	return nil
}

// createMigrationsTable creates the migrations table if it doesn't exist
func (m *Migrator) createMigrationsTable() error {
	var query string
	switch m.dbType {
	case "postgres":
		query = `
			CREATE TABLE IF NOT EXISTS migrations (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				applied_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			)
		`
	case "mysql":
		query = `
			CREATE TABLE IF NOT EXISTS migrations (
				id INT AUTO_INCREMENT PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				applied_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			)
		`
	case "sqlite3":
		query = `
			CREATE TABLE IF NOT EXISTS migrations (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT NOT NULL,
				applied_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			)
		`
	default:
		return fmt.Errorf("unsupported database type: %s", m.dbType)
	}

	_, err := m.db.Exec(query)
	return err
}

// getAppliedMigrations returns the list of applied migrations
func (m *Migrator) getAppliedMigrations() ([]string, error) {
	var migrations []Migration
	if _, err := m.db.Select(&migrations, "SELECT * FROM migrations ORDER BY id"); err != nil {
		return nil, err
	}

	var names []string
	for _, migration := range migrations {
		names = append(names, migration.Name)
	}
	return names, nil
}

// isMigrationApplied checks if a migration has been applied
func (m *Migrator) isMigrationApplied(appliedMigrations []string, migration string) bool {
	for _, applied := range appliedMigrations {
		if applied == migration {
			return true
		}
	}
	return false
}

// applyMigration applies a migration and records it
func (m *Migrator) applyMigration(migration string) error {
	// Execute the migration
	if _, err := m.db.Exec(migration); err != nil {
		return err
	}

	// Record the migration
	_, err := m.db.Exec("INSERT INTO migrations (name) VALUES (?)", migration)
	return err
}
