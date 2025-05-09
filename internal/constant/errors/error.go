package errors

import (
	"net/http"

	"github.com/joomcode/errorx"
)

type ErrorType struct {
	ErrorType  *errorx.Type
	StatusCode int
}

var (
	invalidInput  = errorx.NewNamespace("validation Error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	dbErr         = errorx.NewNamespace("db err")
	alreadyExists = errorx.NewNamespace("data already exists").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	notFound      = errorx.NewNamespace("resource not found").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
)

var (
	ErrInvalidUserInput = errorx.NewType(invalidInput, "invalid input")
	ErrDB               = errorx.NewType(dbErr, "internal server error")
	ErrDataExists       = errorx.NewType(alreadyExists, "data already exists")
	ErrResourceNotFound = errorx.NewType(notFound, "resource not found")
)

var Error = []ErrorType{
	{
		ErrorType:  ErrInvalidUserInput,
		StatusCode: http.StatusBadRequest,
	},
	{
		ErrorType:  ErrDB,
		StatusCode: http.StatusInternalServerError,
	},
	{
		ErrorType:  ErrDataExists,
		StatusCode: http.StatusBadRequest,
	},
	{
		ErrorType:  ErrResourceNotFound,
		StatusCode: http.StatusNotFound,
	},
}
