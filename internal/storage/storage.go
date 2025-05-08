package storage

import (
	"context"
	"url_shortener/internal/constant/model/dto"
)

type URL interface {
	CreateURL(ctx context.Context, urlRequest dto.URLRequest) (*dto.URLResponse, error)
	GetURL(ctx context.Context, shortCode string) (*dto.URLResponse, error)
	GetURLDetails(ctx context.Context, shortCode string) (*dto.URLResponse, error)
	UpdateURL(ctx context.Context, shortCode string, req dto.URLRequest) (*dto.URLResponse, error)
	DeleteURL(ctx context.Context, shortCode string) error
}
