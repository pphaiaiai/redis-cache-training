package logging

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
)

var Logger zerolog.Logger

func NewLogger() {
	env := os.Getenv("STAGE_STATUS")

	if env == "dev" {
		output := zerolog.ConsoleWriter{Out: os.Stdout}
		output.FormatLevel = func(i interface{}) string {
			level := strings.ToUpper(fmt.Sprintf("%s", i))
			switch level {
			case "DEBUG":
				return fmt.Sprintf("| \033[37m%s\033[0m |", level)
			case "INFO":
				return fmt.Sprintf("| \033[32m%s\033[0m |", level)
			case "WARN":
				return fmt.Sprintf("| \033[33m%s\033[0m |", level)
			case "ERROR":
				return fmt.Sprintf("| \033[31m%s\033[0m |", level)
			case "FATAL":
				return fmt.Sprintf("| \033[35m%s\033[0m |", level)
			case "PANIC":
				return fmt.Sprintf("| \033[36m%s\033[0m |", level)
			default:
				return fmt.Sprintf("| %s |", level)
			}
		}
		output.FormatTimestamp = func(i interface{}) string {
			return fmt.Sprintf("\033[37m%s\033[0m", time.Now().Format("15:04:05"))
		}
		output.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("\033[32m%s\033[0m\n", i)
		}
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("\033[35m%s\033[0m: ", i)
		}
		output.FormatFieldValue = func(i interface{}) string {
			return fmt.Sprintf("\033[33m%s\033[0m\n", i)
		}

		output.PartsOrder = []string{
			zerolog.TimestampFieldName,
			zerolog.LevelFieldName,
			zerolog.CallerFieldName,
			zerolog.MessageFieldName,
		}

		Logger = zerolog.New(output).
			Level(zerolog.DebugLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()

	} else {
		Logger = zerolog.New(os.Stdout).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger().
			Hook(zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, message string) {
				fmt.Println()
			}))
	}

	log.Logger = Logger
}

func ZerologMiddleware(logger zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		stop := time.Now()

		latency := stop.Sub(start).String()
		status := c.Response().StatusCode()
		method := c.Method()
		path := c.Path()

		logEvent := logger.Info()
		logEvent.Int("status", status).
			Str("method", method).
			Str("latency", latency).
			Str("path", path).
			Msg("Incoming Request")

		return err
	}
}

type ZerologGormLogger struct {
	SlowThreshold time.Duration
	logger        zerolog.Logger
	logLevel      logger.LogLevel
}

func NewZerologGormLogger(slowThreshold time.Duration, logLevel logger.LogLevel) *ZerologGormLogger {
	return &ZerologGormLogger{
		SlowThreshold: slowThreshold,
		logger:        Logger,
		logLevel:      logLevel,
	}
}

func (z *ZerologGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return &ZerologGormLogger{
		logger:   z.logger,
		logLevel: level,
	}
}

func (z *ZerologGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if z.logLevel >= logger.Info {
		z.logger.Info().Msgf(msg, data...)
	}
}

func (z *ZerologGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if z.logLevel >= logger.Warn {
		z.logger.Warn().Msgf(msg, data...)
	}
}

func (z *ZerologGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if z.logLevel >= logger.Error {
		z.logger.Error().Msgf(msg, data...)
	}
}

func (z *ZerologGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if z.logLevel <= 0 {
		return
	}

	sql, rows := fc()
	elapsed := time.Since(begin)

	switch {
	case err != nil && z.logLevel >= logger.Error:
		z.logger.Error().
			Dur("elapsed", elapsed).
			Int64("rows", rows).
			Err(err).
			Msgf("%s", sql)
	case elapsed > z.SlowThreshold && z.logLevel >= logger.Warn:
		z.logger.Warn().
			Dur("elapsed", elapsed).
			Int64("rows", rows).
			Msgf("%s", sql)
	case z.logLevel >= logger.Info:
		z.logger.Info().
			Dur("elapsed", elapsed).
			Int64("rows", rows).
			Msgf("%s", sql)
	}
}
