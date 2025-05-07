package initiator

import (
	"url_shortener/internal/module"
	"url_shortener/internal/module/url"
	"url_shortener/platform/logger"
)


type Module struct{
	url module.URL
}

func InitModule(persistence Persistence,log logger.Logger) Module {
	return Module{
		url: url.InitURLModule(persistence.url,log.Named("url-module")),
	}
}