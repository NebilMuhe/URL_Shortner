package module

import (
	"context"
	"url_shortener/internal/constant/model/dto"
)

type URL interface {
	CreateURL(ctx context.Context, urlRequest dto.URLRequest) (*dto.URLResponse, error)
}
