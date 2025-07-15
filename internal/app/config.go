package app

import (
	"log/slog"
	"os"
	"submanager/internal/pkg/envzilla"
	"submanager/internal/pkg/postgres"
)

type (
	Config struct {
		Port        string `env:"PORT" default:"localhost"`
		Host        string `env:"HOST" default:"8080"`
		LogLevel    string `env:"LOG_LEVEL" default:"dev"`
		DB          postgres.DBConfig
		LogFilePath string `env:"LOG_FILE_PATH" default:"docs/"`
	}
)

func MustParseConfig() Config {
	if err := envzilla.Loader(".env"); err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	var cfg Config
	if err := envzilla.Parse(&cfg); err != nil {
		slog.Error("Failed to parse configuration", "error", err)
		os.Exit(1)
	}

	return cfg
}
