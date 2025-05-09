package initiator

import (
	"url_shortener/internal/handler/rest"
	"url_shortener/internal/handler/rest/gin/url"
	"url_shortener/platform/logger"
)

type Handler struct {
	urlHandler rest.URL
}

func InitHandler(module Module, log logger.Logger) Handler {
	return Handler{
		urlHandler: url.InitURLHandler(module.url, log.Named("url-handler")),
	}
}
