package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool" // Use this instead of database/sql
)

func ConnectDB(ctx context.Context, connStr string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
