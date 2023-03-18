package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/amyunfei/glassy-sky/internal/admin/infrastructure/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := errors.New("invalid authorization header")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}
		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
