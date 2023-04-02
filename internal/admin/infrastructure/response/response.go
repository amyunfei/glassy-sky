package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Body[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"msg"`
}

func Success[T any](ctx *gin.Context, data T, message string) {
	ctx.JSON(http.StatusOK, Body[T]{
		Code:    0,
		Data:    data,
		Message: message,
	})
}

func UnexpectedError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, Body[*any]{
		Code:    -1,
		Data:    nil,
		Message: message,
	})
}

func RequestError(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, Body[*any]{
		Code:    -1,
		Data:    nil,
		Message: message,
	})
}

type ErrorMessages map[string]string
type Verifiable interface {
	GetValidateMessage() ErrorMessages
}

func ValidationError(ctx *gin.Context, data Verifiable, err error) {
	fmt.Println(err.(validator.ValidationErrors))
	for _, v := range err.(validator.ValidationErrors) {
		if message, exist := data.GetValidateMessage()[v.Field()+"."+v.Tag()]; exist {
			RequestError(ctx, message)
			return
		}
	}
	RequestError(ctx, "parameter Error")
}
