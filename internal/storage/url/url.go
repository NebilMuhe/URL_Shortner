package url

import (
	"context"
	"url_shortener/internal/constant/model/db"
	"url_shortener/internal/constant/model/dto"
	persistencedb "url_shortener/internal/constant/model/persistenceDB"
	"url_shortener/internal/storage"
	"url_shortener/platform/logger"

	"go.uber.org/zap"
)

type URL struct {
	log logger.Logger
	db  persistencedb.PersistenceDB
}

func InitURLPersistence(db persistencedb.PersistenceDB, log logger.Logger) storage.URL {
	return &URL{
		log: log,
		db:  db,
	}
}

func (u *URL) CreateURL(ctx context.Context, urlRequest dto.URLRequest) (*dto.URLResponse, error) {
	url, err := u.db.CreateURL(ctx, db.CreateURLParams{
		OriginalUrl: urlRequest.OriginalURL,
		ShortCode: urlRequest.ShortCode,
	})

	if err != nil {
		u.log.Error(ctx, "failed to create url", zap.Error(err))
		return nil, err
	}
	return &dto.URLResponse{
		ID:          url.ID,
		OriginalURL: url.OriginalUrl,
		ShortCode:   url.ShortCode,
		Count:       url.Count.Int32,
		CreatedAt:   url.CreatedAt,
		UpdatedAt:   url.UpdatedAt,
	}, nil
}
