package response

import "github.com/gin-gonic/gin"



func SuccessResponse(ctx *gin.Context,data any,staus_code int){
	ctx.JSON(
		staus_code,
		Response{
			Ok: true,
			Data: data,
		},
	)
}

func SendErrorResponse(ctx *gin.Context, err *ErrorResponse){
	ctx.AbortWithStatusJSON(err.StausCode,
		Response{
			Ok: false,
			Error: err,
		},
	)
}