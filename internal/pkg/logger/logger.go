package logger

import (
	"context"
	"log/slog"
	"os"
)

const (
	Debug = "debug"
	Prod  = "prod"
	Dev   = "dev"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	With(args ...any) Logger
	Log(ctx context.Context, level slog.Level, msg string, args ...any)
}

type slogLogger struct {
	*slog.Logger
}

func (l *slogLogger) Debug(msg string, args ...any) {
	l.Logger.Debug(msg, args...)
}

func (l *slogLogger) Info(msg string, args ...any) {
	l.Logger.Info(msg, args...)
}

func (l *slogLogger) Warn(msg string, args ...any) {
	l.Logger.Warn(msg, args...)
}

func (l *slogLogger) Error(msg string, args ...any) {
	l.Logger.Error(msg, args...)
}

func (l *slogLogger) With(args ...any) Logger {
	return &slogLogger{
		Logger: l.Logger.With(args...),
	}
}

func (l *slogLogger) Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	l.Logger.Log(ctx, level, msg, args...)
}

func New(level string) Logger {
	var slogger *slog.Logger
	switch level {
	case Debug:
		slogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case Prod:
		slogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	case Dev:
		slogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	default:
		slogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}
	return &slogLogger{slogger}
}
