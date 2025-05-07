package initiator

import (
	"url_shortener/internal/handler"
	"url_shortener/internal/handler/url"
	"url_shortener/platform/logger"
)

type Handler struct {
	urlHandler handler.URL
}

func InitHandler(module Module, log logger.Logger) Handler {
	return Handler{
		urlHandler: url.InitURLHandler(module.url, log.Named("url-handler")),
	}
}
