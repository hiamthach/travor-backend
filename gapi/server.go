package gapi

import (
	"context"
	"log"

	"github.com/travor-backend/dto"
	"github.com/travor-backend/model"
	"github.com/travor-backend/pb"
	"github.com/travor-backend/util"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type GRPCServer struct {
	config     util.Config
	store      *gorm.DB
	tokenMaker util.Maker
	pb.UnimplementedAuthServiceServer
}

func NewGRPCServer(store *gorm.DB, config util.Config) (*GRPCServer, error) {
	tokenMaker, err := util.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("Can not create token maker: ", err)
	}
	return &GRPCServer{store: store, config: config, tokenMaker: tokenMaker}, nil
}

func (server *GRPCServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user model.User

	err := server.store.Table("users").Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return nil, err
	}

	if err := util.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return nil, err
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration, server.config.AccessTokenPrivateKey)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.RefreshTokenDuration, server.config.RefreshTokenPrivateKey)
	if err != nil {
		return nil, err
	}

	session := dto.UserSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}

	if err := server.store.Create(&session).Error; err != nil {
		return nil, err
	}

	rsp := pb.LoginResponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}

	return &rsp, nil
}

func (server *GRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Username:       req.Username,
		Email:          req.Email,
		Phone:          req.Phone,
		FullName:       req.FullName,
		HashedPassword: hashedPassword,
		Status:         true,
		Role:           1,
	}

	if err := server.store.Create(&user).Error; err != nil {
		return nil, err
	}

	return convertUser(user), nil
}
