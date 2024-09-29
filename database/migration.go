package database

import (
	"database/sql"
	"fmt"
	"os"
	"redis-cache-training/logging"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate() error {
	logger := logging.Logger.With().Str("method", "Migrate").Logger()

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get db driver")
	}
	m, err := migrate.NewWithDatabaseInstance(os.Getenv("SQL_SOURCE_PATH"), "postgres", driver)
	if err != nil {
		logger.Fatal().Err(err)
		return err
	}
	//nolint:golint // make it catch error once no changes
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			logger.Info().Msg("No changes, database schema is up to date")
			return nil
		}
		logger.Fatal().Err(err)
		return err
	} // or m.Step(2) if you want to explicitly set the number of migrations to run
	logger.Info().Msg("Database schema migrated successfully")
	return nil
}
