package middleware

import (
	"net/http"
	"url_shortener/internal/constant/errors"
	"url_shortener/internal/constant/model/response"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/joomcode/errorx"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			e := ctx.Errors[0]
			err := e.Unwrap()
			errRes := CastErrorResponse(err)
			if errRes != nil {
				response.SendErrorResponse(ctx, errRes)
				return
			}

			response.SendErrorResponse(ctx, &response.ErrorResponse{
				StausCode: http.StatusInternalServerError,
				Message:   "internal server error",
			})
		}
	}
}

func CastErrorResponse(err error) *response.ErrorResponse {
	castedError := errorx.Cast(err)
	if castedError == nil {
		return nil
	}
	if code, ok := errors.ErrorMap[castedError.Type()]; ok {
		return &response.ErrorResponse{
			StausCode:  code,
			Message:    castedError.Message(),
			FieldError: FieldErrors(castedError.Cause()),
		}
	}
	return nil
}

func FieldErrors(err error) []response.FieldError {
	var fieldErrors []response.FieldError

	if data, ok := err.(validation.Errors); ok {
		for k, v := range data {
			fieldErrors = append(fieldErrors, response.FieldError{
				Name:        k,
				Description: v.Error(),
			})
		}

		return fieldErrors
	}

	return nil

}
