package url

import (
	"net/http"
	"url_shortener/internal/glue/routes"
	"url_shortener/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRoute(group *gin.RouterGroup, url handler.URL) {
	urlRoutes := []routes.Routes{
		{
			Path:        "/url",
			Method:      http.MethodPost,
			Handler:     url.CreateURL,
			Middlewares: []gin.HandlerFunc{},
		},
	}

	routes.RegisterRoutes(group, urlRoutes)
}
