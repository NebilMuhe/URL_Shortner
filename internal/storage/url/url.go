package url

import (
	"context"
	"url_shortener/internal/constant/errors"
	"url_shortener/internal/constant/model/db"
	"url_shortener/internal/constant/model/dto"
	persistencedb "url_shortener/internal/constant/model/persistenceDB"
	"url_shortener/internal/storage"
	"url_shortener/platform/logger"

	"github.com/joomcode/errorx"
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
		ShortCode:   urlRequest.ShortCode,
	})

	if err != nil {
		if errorx.IsDuplicate(err) {
			err = errors.ErrDataExists.Wrap(err, "url already exists")
			u.log.Info(ctx, "url already exist", zap.Error(err))
			return nil, err
		}
		err = errors.ErrDB.Wrap(err, "error while creating url")
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

func (u *URL) GetURL(ctx context.Context, shortCode string) (*dto.URLResponse, error) {
	url, err := u.db.UpdateCount(ctx, shortCode)
	if err != nil {
		if errorx.IsNotFound(err) {
			err = errors.ErrResourceNotFound.Wrap(err, "url not exists")
			u.log.Info(ctx, "url not found", zap.Error(err))
			return nil, err
		}
		err = errors.ErrDB.Wrap(err, "error while getting url")
		u.log.Error(ctx, "failed to get url", zap.Error(err))
		return nil, err
	}
	return &dto.URLResponse{
		ID:          url.ID,
		OriginalURL: url.OriginalUrl,
		ShortCode:   url.ShortCode,
		Count:       url.Count.Int32,
		CreatedAt:   url.CreatedAt,
		UpdatedAt:   url.UpdatedAt,
		DeletedAt:   url.DeletedAt,
	}, nil
}

func (u *URL) GetURLDetails(ctx context.Context, shortCode string) (*dto.URLResponse, error) {
	url, err := u.db.GetURLByShortCode(ctx, shortCode)
	if err != nil {
		if errorx.IsNotFound(err) {
			err = errors.ErrResourceNotFound.Wrap(err, "url not exists")
			u.log.Info(ctx, "url not found", zap.Error(err))
			return nil, err
		}
		err = errors.ErrDB.Wrap(err, "errpr while getting url details")
		u.log.Error(ctx, "failed to get url details", zap.Error(err))
		return nil, err
	}

	return &dto.URLResponse{
		ID:          url.ID,
		OriginalURL: url.OriginalUrl,
		ShortCode:   url.ShortCode,
		Count:       url.Count.Int32,
		CreatedAt:   url.CreatedAt,
		UpdatedAt:   url.UpdatedAt,
		DeletedAt:   url.DeletedAt,
	}, nil
}

func (u *URL) UpdateURL(ctx context.Context, shortCode string, req dto.URLRequest) (*dto.URLResponse, error) {
	url, err := u.db.UpdateURL(ctx, db.UpdateURLParams{
		OriginalUrl: req.OriginalURL,
		ShortCode:   shortCode,
	})
	if err != nil {
		if errorx.IsNotFound(err) {
			err = errors.ErrResourceNotFound.Wrap(err, "url not exists")
			u.log.Info(ctx, "url not found", zap.Error(err))
			return nil, err
		}
		if errorx.IsDuplicate(err) {
			err = errors.ErrDataExists.Wrap(err, "url already exists")
			u.log.Info(ctx, "url already exist", zap.Error(err))
			return nil, err
		}
		err = errors.ErrDB.Wrap(err, "error while updating url")
		u.log.Error(ctx, "failed to update url", zap.Error(err))
		return nil, err
	}
	return &dto.URLResponse{
		ID:          url.ID,
		OriginalURL: url.OriginalUrl,
		ShortCode:   url.ShortCode,
		Count:       url.Count.Int32,
		CreatedAt:   url.CreatedAt,
		UpdatedAt:   url.UpdatedAt,
		DeletedAt:   url.DeletedAt,
	}, nil
}

func (u *URL) DeleteURL(ctx context.Context, shortCode string) error {
	if err := u.db.DeleteURL(ctx, shortCode); err != nil {
		if errorx.IsNotFound(err) {
			err = errors.ErrResourceNotFound.Wrap(err, "url not exists")
			u.log.Info(ctx, "url not found", zap.Error(err))
			return err
		}
		err = errors.ErrDB.Wrap(err, "error while deleting url")
		u.log.Error(ctx, "failed to delete url", zap.Error(err))
		return err
	}
	return nil
}
