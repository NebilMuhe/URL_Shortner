package handler

import "github.com/gin-gonic/gin"

type URL interface {
	CreateURL(ctx *gin.Context)
}
