package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type API struct {
	Pool *pgxpool.Pool
}

type DBConfig struct {
	Host           string        `env:"POSTGRES_HOST"`
	Port           string        `env:"POSTGRES_PORT"`
	Name           string        `env:"POSTGRES_DB"`
	UserName       string        `env:"POSTGRES_USER"`
	Password       string        `env:"POSTGRES_PASSWORD"`
	MaxConnections int32         `env:"POSTGRES_MAX_CONNS" default:"10"`
	ConnTimeout    time.Duration `env:"POSTGRES_CONN_TIMEOUT" default:"5s"`
}

// Connect connects to the PostgreSQL database using the provided configuration.
func Connect(ctx context.Context, dbCfg DBConfig) (*API, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbCfg.UserName, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig: %w", err)
	}

	cfg.MaxConns = dbCfg.MaxConnections
	cfg.MaxConnLifetime = time.Hour
	cfg.HealthCheckPeriod = time.Minute

	ctx, cancel := context.WithTimeout(ctx, dbCfg.ConnTimeout)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &API{
		Pool: pool,
	}, nil
}

// Close closes the database connection gracefully.
func (a *API) Close() {
	a.Pool.Close()
}
