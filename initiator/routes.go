package initiator

import (
	"url_shortener/internal/glue/routes/url"

	"github.com/gin-gonic/gin"
)

func InitRoute(group *gin.RouterGroup, handler Handler) {
	url.InitRoute(group, handler.urlHandler)
}
