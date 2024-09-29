package utils

import (
	"fmt"
	"os"

	errorsConstant "redis-cache-training/errors"

	"github.com/pkg/errors"
)

func ConnectionURLBuilder(str string) (string, error) {
	var url string

	switch str {
	case "postgres":
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSL"),
		)
	case "redis":
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		)
	case "fiber":
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		// Return error message.
		return "", errors.Errorf(errorsConstant.ERROR_CODE_CONNECTION_NOT_SUPPORTED)
	}

	return url, nil
}
