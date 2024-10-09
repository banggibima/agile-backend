package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/banggibima/agile-backend/config"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Client(config *config.Config, logger *logrus.Logger) (*sql.DB, error) {
	driver := config.Postgres.Driver
	url := config.Postgres.URL
	sslmode := config.Postgres.SSLMode

	client, err := sql.Open(driver, url+"?sslmode="+sslmode)
	if err != nil {
		return nil, err
	}

	if err := Connect(client); err != nil {
		return nil, err
	}

	if err := Migration(client, logger); err != nil {
		return nil, err
	}

	return client, nil
}

func Connect(client *sql.DB) error {
	err := client.Ping()
	if err != nil {
		return err
	}

	return nil
}

func Migration(db *sql.DB, logger *logrus.Logger) error {
	query := "CREATE TABLE IF NOT EXISTS migrations (	id SERIAL PRIMARY KEY, name VARCHAR(255),	applied_at TIMESTAMPTZ DEFAULT NOW())"

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %v", err)
	}

	files, err := os.ReadDir("migration")
	if err != nil {
		return fmt.Errorf("failed to read migration files: %v", err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		count := 0
		query = "SELECT COUNT(*) FROM migrations WHERE name = $1"

		err := db.QueryRow(query, file.Name()).Scan(&count)
		if err != nil {
			return fmt.Errorf("failed to check migration: %v", err)
		}

		if count > 0 {
			continue
		}

		path := filepath.Join("migration", file.Name())
		sql, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read migration file: %v", err)
		}

		_, err = db.Exec(string(sql))
		if err != nil {
			return fmt.Errorf("failed to execute migration: %v", err)
		}

		query = "INSERT INTO migrations (name) VALUES ($1)"

		_, err = db.Exec(query, file.Name())
		if err != nil {
			return fmt.Errorf("failed to record migration: %v", err)
		}

		logger.Infof("migration applied: %s", file.Name())
	}

	return nil
}

func Seed(db *sql.DB, logger *logrus.Logger) error {
	files, err := os.ReadDir("seed")
	if err != nil {
		return fmt.Errorf("failed to read seed files: %v", err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		path := filepath.Join("seed", file.Name())
		sql, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read seed file: %v", err)
		}

		_, err = db.Exec(string(sql))
		if err != nil {
			return fmt.Errorf("failed to execute seed: %v", err)
		}

		logger.Infof("seed applied: %s", file.Name())
	}

	return nil
}
