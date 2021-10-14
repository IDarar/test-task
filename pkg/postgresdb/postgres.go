package postgresdb

import (
	"context"
	"fmt"
	"os"

	"github.com/IDarar/test-task/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresDB(cfg config.Config) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(cfg.Postgres.URL)
	if err != nil {
		return nil, err
	}

	config.ConnConfig.PreferSimpleProtocol = true

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to connect to database:", err)
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		fmt.Printf("postgres does not respond: %v\n", err)
		return nil, err
	}

	return pool, nil
}
