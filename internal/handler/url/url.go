package url

import (
	"context"
	"net/http"
	"time"
	"url_shortener/internal/constant/model/dto"
	"url_shortener/internal/handler"
	"url_shortener/internal/module"
	"url_shortener/platform/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type URL struct {
	log     logger.Logger
	module  module.URL
	timeout time.Duration
}

func InitURLHandler(urlModule module.URL, log logger.Logger) handler.URL {
	return &URL{
		log:    log,
		module: urlModule,
	}
}

func (u *URL) CreateURL(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()
	var createUrl dto.URLRequest
	if err := ctx.ShouldBind(&createUrl); err != nil {
		u.log.Info(ctx, "failed to bind create url request", zap.Error(err))
		ctx.Error(err)
		return
	}

	response, err := u.module.CreateURL(c, createUrl)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (u *URL) GetURL(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	short_code := ctx.Param("short_code")

	res, err := u.module.GetURL(c, short_code)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Redirect(http.StatusOK, res.OriginalURL)
}

func (u *URL) GetURLDetails(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	short_code := ctx.Param("short_code")
	res, err := u.module.GetURLDetails(c, short_code)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (u *URL) UpdateURL(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	var urlRequest dto.URLRequest

	short_code := ctx.Param("short_code")

	if err := ctx.ShouldBind(&urlRequest); err != nil {
		u.log.Info(c, "failed to bind request", zap.Error(err))
		ctx.Error(err)
		return
	}

	res, err := u.module.UpdateURL(c, short_code, urlRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (u *URL) DeleteURL(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	short_code := ctx.Param("short_code")
	if err := u.module.DeleteURL(c, short_code); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
