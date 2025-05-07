package persistencedb

import (
	"url_shortener/internal/constant/model/db"
	"url_shortener/platform/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PersistenceDB struct {
	*db.Queries
	log  logger.Logger
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool, log logger.Logger) PersistenceDB {
	return PersistenceDB{
		log:     log,
		pool:    pool,
		Queries: db.New(pool),
	}
}
