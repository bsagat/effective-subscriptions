package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	httpserver "submanager/internal/adapters/http"
	"submanager/internal/adapters/repo"
	"submanager/internal/core/service"
	"submanager/internal/pkg/logger"
	"submanager/internal/pkg/postgres"
	"syscall"
)

type App struct {
	httpServer *httpserver.API
	postgresDB *postgres.API

	log logger.Logger
}

// Setup application with adapters and logger
func New(cfg Config, log logger.Logger) *App {
	ctx := context.Background()

	log.Info("Connecting to database...")
	postgresDB, err := postgres.Connect(ctx, cfg.DB)
	if err != nil {
		log.Error("Failed to connect postgres server", "error", err)
		os.Exit(1)
	}
	log.Info("Database connection estabilished...")

	subsRepo := repo.NewSubsRepo(postgresDB.Pool)
	subsService := service.NewSubsService(subsRepo, log)
	server := httpserver.New(cfg.Host, cfg.Port, subsService, log)

	return &App{
		httpServer: server,
		postgresDB: postgresDB,
		log:        log,
	}
}

// Start the application, which includes starting the HTTP server and handling graceful shutdown
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

// CleanUp closes the HTTP server and database connection gracefully
func (a *App) CleanUp() {
	if err := a.httpServer.Close(); err != nil {
		a.log.Error("Failed to close server...")
	}

	a.postgresDB.Close()
}
