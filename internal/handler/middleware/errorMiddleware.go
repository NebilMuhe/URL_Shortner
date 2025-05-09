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
	for _, e := range errors.Error {
		if errorx.IsOfType(err, e.ErrorType) {
			er := errorx.Cast(err)
			if err == nil {
				return &response.ErrorResponse{
					StausCode: http.StatusInternalServerError,
					Message:   "internal server error",
				}
			}
			return &response.ErrorResponse{
				StausCode:  e.StatusCode,
				Message:    er.Message(),
				FieldError: FieldErrors(er.Cause()),
			}
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
