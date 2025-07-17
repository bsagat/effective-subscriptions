package main

import (
	"submanager/internal/app"
	"submanager/internal/pkg/logger"
)

func main() {
	cfg := app.MustParseConfig()

	log := logger.New(cfg.LogLevel)
	log.Info("Logger setup finished")

	application := app.New(cfg, log)
	application.Start()

	log.Info("Application closed successfully")
}
