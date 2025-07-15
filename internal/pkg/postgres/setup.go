package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type API struct {
	DB *pgx.Conn
}

type DBConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	Name     string `env:"POSTGRES_DB"`
	UserName string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
}

func Connect(dbCfg DBConfig) (*API, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbCfg.UserName, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
	pool, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &API{
		DB: pool,
	}, nil
}

func (a *API) Close(ctx context.Context) error {
	return a.DB.Close(ctx)
}
