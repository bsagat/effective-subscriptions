package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	httpserver "submanager/internal/adapters/http"
	"submanager/internal/pkg/postgres"
	"syscall"
)

type App struct {
	httpServer *httpserver.API
	postgresDB *postgres.API

	log *slog.Logger
}

func New(cfg Config, log *slog.Logger) *App {
	server := httpserver.New(cfg.Host, cfg.Port, log)

	log.Info("Connecting to database...")
	postgresDB, err := postgres.Connect(cfg.DB)
	if err != nil {
		log.Error("Failed to connect postgres server", "error", err)
		os.Exit(1)
	}
	log.Info("Database connection estabilished...")

	return &App{
		httpServer: server,
		postgresDB: postgresDB,
		log:        log,
	}
}

func (a *App) Start() {
	go a.httpServer.StartServer()
	defer a.CleanUp()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	s := <-stop
	a.log.Info(fmt.Sprintf("Catched %s signal!!!", s.String()))
	a.log.Info("Shutting down server")
}

func (a *App) CleanUp() {
	if err := a.httpServer.Close(); err != nil {
		a.log.Error("Failed to close server...")
	}

	if err := a.postgresDB.Close(context.Background()); err != nil {
		a.log.Error("Failed to close database...")
	}
}
