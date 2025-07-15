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

func New(logFilePath string, level string) *slog.Logger {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("Failed to open/create log file", "error", err)
		os.Exit(1)
	}

	var logger *slog.Logger
	switch level {
	case debug:
		logger = slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case prod:
		logger = slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	case dev:
		logger = slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}
	return logger
}
