package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vgekko/ani-go/config"
	"golang.org/x/exp/slog"
	"log"
	"time"
)

func NewPool(cfg config.Postgres) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool

	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Database, cfg.Password, cfg.SSLMode)

	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("postgresql.NewPool: %w", err)
	}

	poolConfig.MaxConns = int32(cfg.MaxPoolSize)

	connAttempts := cfg.ConnAttempts
	connTimeout := cfg.ConnTimeout

	for connAttempts > 0 {
		pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}

		log.Printf("postgresql trying to connect, attempts left: %d", connAttempts)

		// sleep until connection timeout expires
		time.Sleep(connTimeout)

		connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgresql.NewPool.connAttempts == 0: %w", err)
	}

	slog.Info("connected to postgresql database: ", cfg.Database)
	return pool, nil
}
