package gapi

import (
	"context"

	"github.com/travor-backend/constant"
	"github.com/travor-backend/db"

	"github.com/travor-backend/interceptor"
	"github.com/travor-backend/model"
	"github.com/travor-backend/pb"
	"github.com/travor-backend/util"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type AuthServer struct {
	config util.Config
	store  *gorm.DB
	pb.UnimplementedAuthServiceServer
}

func NewAuthServer(config util.Config) (*AuthServer, error) {
	userdb, err := db.Connect(config.GetDBSource(constant.USER_DB))
	if err != nil {
		return nil, err
	}

	return &AuthServer{store: userdb, config: config}, nil
}

func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user model.User

	err := server.store.Table("users").Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return nil, err
	}

	if err := util.CheckPassword(req.Password, user.HashedPassword); err != nil {
		return nil, err
	}

	accessToken, accessPayload, err := util.CreateToken(user.Username, server.config.AccessTokenDuration, server.config.AccessTokenPrivateKey, user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshPayload, err := util.CreateToken(user.Username, server.config.RefreshTokenDuration, server.config.RefreshTokenPrivateKey, user.Role)
	if err != nil {
		return nil, err
	}

	session := model.UserSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}

	if err := server.store.Table("sessions").Create(&session).Error; err != nil {
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

func (server *AuthServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
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

func (server *AuthServer) RenewToken(ctx context.Context, req *pb.RenewTokenRequest) (*pb.RenewTokenResponse, error) {
	refreshToken, err := util.VerifyToken(req.RefreshToken, server.config.RefreshTokenPublicKey)
	if err != nil {
		return nil, err
	}

	var session model.UserSessionParams
	if err := server.store.Table("sessions").Where("id = ?", refreshToken.ID).First(&session).Error; err != nil {
		return nil, err
	}

	if session.IsBlocked {
		return nil, err
	}

	if session.RefreshToken != req.RefreshToken {
		return nil, err
	}

	var user model.User
	if err := server.store.Table("sessions").Where("username = ?", session.Username).First(&user).Error; err != nil {
		return nil, err
	}

	accessToken, accessPayload, err := util.CreateToken(session.Username, server.config.AccessTokenDuration, server.config.AccessTokenPrivateKey, user.Role)
	if err != nil {
		return nil, err
	}

	return &pb.RenewTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiredAt),
	}, nil
}

func (server *AuthServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	var user model.User

	err := server.store.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}

func (server *AuthServer) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	c, err := interceptor.AuthInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)

	var user model.User

	if err := server.store.Where("username = ?", payload.Username).First(&user).Error; err != nil {
		return nil, err
	}

	if err := util.CheckPassword(req.OldPassword, user.HashedPassword); err != nil {
		return nil, err
	}

	hashedPassword, err := util.HashPassword(req.NewPassword)
	if err != nil {
		return nil, err
	}

	if err := server.store.Model(&user).Update("hashed_password", hashedPassword).Error; err != nil {
		return nil, err
	}

	return &pb.ChangePasswordResponse{
		Success: true,
	}, nil
}
