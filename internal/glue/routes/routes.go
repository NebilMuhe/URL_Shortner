package routes

import "github.com/gin-gonic/gin"

type Routes struct {
	Path        string
	Method      string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

func RegisterRoutes(group *gin.RouterGroup, routes []Routes) {
	var handler []gin.HandlerFunc

	for _, route := range routes {
		handler = append(handler, route.Middlewares...)
		handler = append(handler, route.Handler)

		group.Handle(route.Method, route.Path, handler...)
	}
}
