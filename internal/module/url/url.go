package url

import (
	"context"
	"url_shortener/internal/constant/errors"
	"url_shortener/internal/constant/model/dto"
	"url_shortener/internal/module"
	"url_shortener/internal/storage"
	"url_shortener/platform/logger"
	"url_shortener/platform/utils"

	"go.uber.org/zap"
)

type URL struct {
	log            logger.Logger
	urlPersistence storage.URL
}

func InitURLModule(urlStorage storage.URL, log logger.Logger) module.URL {
	return &URL{
		log:            log,
		urlPersistence: urlStorage,
	}
}

func (u URL) CreateURL(ctx context.Context, urlRequest dto.URLRequest) (*dto.URLResponse, error) {
	if err := urlRequest.Validate(); err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		u.log.Info(ctx, "invalid url request", zap.Error(err))
		return nil, err
	}

	shortCode := utils.GenerateRandomString(5)
	urlRequest.ShortCode = shortCode

	res, err := u.urlPersistence.CreateURL(ctx, urlRequest)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *URL) GetURL(ctx context.Context, shortCode string) (*dto.URLResponse, error) {
	urlResponse, err := u.urlPersistence.GetURL(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	return urlResponse, nil
}

func (u *URL) GetURLDetails(ctx context.Context, shortCode string) (*dto.URLResponse, error) {
	res, err := u.urlPersistence.GetURLDetails(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *URL) UpdateURL(ctx context.Context, shortCode string, req dto.URLRequest) (*dto.URLResponse, error) {
	if err := req.Validate(); err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		u.log.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}
	url, err := u.urlPersistence.UpdateURL(ctx, shortCode, req)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (u *URL) DeleteURL(ctx context.Context, shortCode string) error {
	if err := u.urlPersistence.DeleteURL(ctx, shortCode); err != nil {
		return err
	}

	return nil
}
