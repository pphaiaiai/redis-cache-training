package database

import (
	"redis-cache-training/logging"
	"redis-cache-training/utils"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDBConnection() (*gorm.DB, error) {
	zerolog := logging.Logger.With().Str("method", "InitDBConnection").Logger()

	postgresDns, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		return nil, err
	}

	db, err = gorm.Open(postgres.Open(postgresDns), &gorm.Config{
		Logger: logging.NewZerologGormLogger(time.Second, logger.Info),
	})

	if err != nil {
		zerolog.Fatal().Err(err).Msg("Failed to connect to database")
	}

	zerolog.Info().Msg("Database successfully initialized")
	return db, nil
}
