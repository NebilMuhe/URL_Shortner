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
		{
			Path:        "/url/:short_code",
			Method:      http.MethodGet,
			Handler:     url.GetURL,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Path:        "/url_details/:short_code",
			Method:      http.MethodGet,
			Handler:     url.GetURLDetails,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Path:        "/url/:short_code",
			Method:      http.MethodPatch,
			Handler:     url.UpdateURL,
			Middlewares: []gin.HandlerFunc{},
		},
		{
			Path:        "/url/:short_code",
			Method:      http.MethodDelete,
			Handler:     url.DeleteURL,
			Middlewares: []gin.HandlerFunc{},
		},
	}

	routes.RegisterRoutes(group, urlRoutes)
}
