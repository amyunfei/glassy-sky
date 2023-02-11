package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"msg"`
}

func Success(ctx *gin.Context, data any, message string) {
	ctx.JSON(http.StatusOK, Body{
		Code:    0,
		Data:    data,
		Message: message,
	})
}

func UnexpectedError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, Body{
		Code:    -1,
		Data:    nil,
		Message: message,
	})
}

func ValidationError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnprocessableEntity, Body{
		Code:    -1,
		Data:    nil,
		Message: message,
	})
}
