package main

import (
	"log"
	"redis-cache-training/database"
	"redis-cache-training/internal/routes"
	"redis-cache-training/logging"
	"redis-cache-training/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	app := fiber.New()

	// Initialize Logger
	logging.NewLogger()
	logger := logging.Logger.With().Str("method", "main").Logger()
	app.Use(logging.ZerologMiddleware(logger))

	db, err := database.InitDBConnection()
	if err != nil {
		logger.Fatal().Err(err).Msg("Error initializing database connection")
	}

	errMigrateDB := database.Migrate()
	if errMigrateDB != nil {
		logger.Fatal().Err(err).Msg("Error migrating database")
	}

	routes.SetupProductRoutes(db, app)

	fiberConnURL, _ := utils.ConnectionURLBuilder("fiber")
	logger.Info().Msgf("Starting server on %s", fiberConnURL)

	logger.Fatal().Err(app.Listen(fiberConnURL)).Msg("Error starting server")
}
