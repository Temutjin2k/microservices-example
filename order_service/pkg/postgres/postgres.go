package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Dsn          string `env:"POSTGRES_DSN,required"`
	MaxOpenConns int32  `env:"POSTGRES_MAX_OPEN_CONN" envDefault:"25"`
	MaxIdleConns int    `env:"POSTGRES_MAX_IDLE_CONN" envDefault:"25"`
	MaxIdleTime  string `env:"POSTGRES_MAX_IDLE_TIME" envDefault:"15m"`
}

func New(ctx context.Context, config Config) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(config.Dsn)
	if err != nil {
		return nil, err
	}

	dbConfig.MaxConns = config.MaxOpenConns
	dbConfig.MinConns = 2

	// Use the time.ParseDuration() function to convert the idle timeout duration string
	// to a time.Duration type.
	duration, err := time.ParseDuration(config.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	dbConfig.MaxConnIdleTime = duration

	dbpool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	return dbpool, nil
}
