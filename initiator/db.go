package initiator

import (
	"context"
	"time"
	"url_shortener/platform/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Options struct {
	Url             string
	MaxConnIdleTime time.Duration
}

func InitDB(ctx context.Context, options Options, log logger.Logger) *pgxpool.Pool{
	config, err := pgxpool.ParseConfig(options.Url)
	if err != nil {
		log.Fatal(ctx, "failed to parse config", zap.Error(err))
	}

	config.ConnConfig.Tracer = log.Named("pgx")
	config.MaxConnIdleTime = options.MaxConnIdleTime

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal(ctx, "failed to create a new pool", zap.Error(err))
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal(ctx, "failed to ping database", zap.Error(err))
	}

	return pool
}
