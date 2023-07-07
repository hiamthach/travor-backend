package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/travor-backend/constant"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		config, err := util.LoadConfig(".")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		authorizationHeader := ctx.GetHeader(constant.AUTHORIZATION_HEADER_KEY)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		fields := strings.Fields(authorizationHeader)

		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != constant.AUTHORIZATION_TYPE_BEARER {
			err := errors.New("authorization type is not Bearer")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		accessToken := fields[1]

		payload, err := util.VerifyToken(accessToken, config.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		ctx.Set(constant.AUTHORIZATION_PAYLOAD_KEY, payload)
		ctx.Set(constant.AUTHORIZATION_ROLE, 1)
		ctx.Next()
	}
}

func AdminMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := ctx.MustGet(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)

		var role int

		if err := db.Table("users").Select("role").Where("username = ?", payload.Username).Scan(&role).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		if role != constant.ADMIN_ROLE {
			err := errors.New("no permission")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		ctx.Set(constant.AUTHORIZATION_ROLE, role)
		ctx.Next()
	}
}
