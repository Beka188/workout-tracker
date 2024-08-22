package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func ConnectDB() (*pgxpool.Pool, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.User, config.Password, config.Host, config.Port, config.DBname, config.SSLMode)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pool, err = pgxpool.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return pool, err
}

func CloseDB() {
	if pool != nil {
		pool.Close()
	}
}
