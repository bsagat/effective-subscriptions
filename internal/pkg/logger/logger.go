package logger

import (
	"log/slog"
	"os"
)

const (
	debug = "debug"
	prod  = "prod"
	dev   = "dev"
)

// New creates a new logger instance with the specified log file path and level.
func New(level string) *slog.Logger {
	var logger *slog.Logger
	switch level {
	case debug:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case prod:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	case dev:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}
	return logger
}
