package interceptor

import (
	"context"
	"errors"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/travor-backend/constant"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

func AuthInterceptor(ctx context.Context) (context.Context, error) {
	config, err := util.LoadConfig(".")
	if err != nil {
		return nil, err
	}

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	payload, err := util.VerifyToken(token, config.AccessTokenPublicKey)
	if err != nil {
		return nil, err
	}

	newCtx := context.WithValue(ctx, constant.AUTHORIZATION_PAYLOAD_KEY, payload)
	return newCtx, nil
}

func AdminInterceptor(ctx context.Context, db *gorm.DB) (context.Context, error) {
	c, err := AuthInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)

	var role int

	if err := db.Table("users").Select("role").Where("username = ?", payload.Username).Scan(&role).Error; err != nil {
		return nil, err
	}

	if role != constant.ADMIN_ROLE {
		err := errors.New("no permission")
		return nil, err
	}

	return c, nil
}
