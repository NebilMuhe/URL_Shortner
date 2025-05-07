package initiator

import (
	persistencedb "url_shortener/internal/constant/model/persistenceDB"
	"url_shortener/internal/storage"
	"url_shortener/internal/storage/url"
	"url_shortener/platform/logger"
)

type Persistence struct {
	url storage.URL
}

func InitPersistence(db persistencedb.PersistenceDB, log logger.Logger) Persistence {
	return Persistence{
		url: url.InitURLPersistence(db, log.Named("url-persistence")),
	}
}
