package handler

import "github.com/gin-gonic/gin"

type URL interface {
	CreateURL(ctx *gin.Context)
	GetURL(ctx *gin.Context)
	GetURLDetails(ctx *gin.Context)
	UpdateURL(ctx *gin.Context)
}
