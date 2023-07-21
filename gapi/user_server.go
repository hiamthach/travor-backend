package gapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/travor-backend/constant"
	"github.com/travor-backend/db"
	"github.com/travor-backend/interceptor"
	"github.com/travor-backend/model"
	"github.com/travor-backend/pb"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

type UserServer struct {
	cache  util.RedisUtil
	config util.Config
	store  *gorm.DB
	pb.UnimplementedUserServiceServer
}

func NewUserServer(config util.Config, cache util.RedisUtil) (*UserServer, error) {
	userdb, err := db.Connect(config.GetDBSource(constant.USER_DB))
	if err != nil {
		return nil, err
	}

	return &UserServer{store: userdb, config: config, cache: cache}, nil
}

func (server *UserServer) GetUserStats(ctx context.Context, req *pb.GetUserStatsRequest) (*pb.GetUserStatsResponse, error) {
	redisKey := fmt.Sprintf("%s:all", constant.USER_REDIS)

	cachedResponse, err := server.cache.Get(ctx, redisKey)
	if err == nil {
		var response *pb.GetUserStatsResponse
		err = json.Unmarshal([]byte(cachedResponse), &response)
		if err != nil {
			return nil, err
		}
		return response, err
	}

	var users []*pb.UserStats
	if err := server.store.Table("users").Select("username", "email", "full_name", "role").Find(&users).Error; err != nil {
		return nil, err
	}

	res := &pb.GetUserStatsResponse{
		Users: users,
	}

	// Serialize the response and store it in Redis cache for future use
	serializedResponse, err := json.Marshal(res)
	if err != nil {
		log.Print(err)
	}

	if err = server.cache.Set(ctx, redisKey, serializedResponse, time.Hour); err != nil {
		log.Print(err)
	}

	return res, nil
}

func (server *UserServer) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	// Add auth interceptor
	c, err := interceptor.AuthInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	if payload.RoleID != constant.ADMIN_ROLE && payload.Username != req.Username {
		return nil, errors.New("unauthorized")
	}

	// Get from redis
	redisKey := fmt.Sprintf("%s:%s", constant.USER_REDIS, req.Username)

	cachedResponse, err := server.cache.Get(ctx, redisKey)
	if err == nil {
		// If the response is found in the cache, deserialize it and return
		var response *pb.GetUserInfoResponse
		err = json.Unmarshal([]byte(cachedResponse), &response)
		if err != nil {
			return nil, err
		}
		return response, err
	}

	var user model.User

	if err := server.store.Table("users").Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, err
	}

	res := &pb.GetUserInfoResponse{
		User: convertUser(user),
	}

	// Serialize the response and store it in Redis cache for future use
	serializedResponse, err := json.Marshal(res)
	if err != nil {
		log.Print(err)
	}

	if err = server.cache.Set(ctx, redisKey, serializedResponse, time.Hour); err != nil {
		log.Print(err)
	}

	return res, nil
}

func (server *UserServer) UpdateUserInfo(ctx context.Context, req *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoResponse, error) {
	// Add auth interceptor
	c, err := interceptor.AuthInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	if payload.RoleID != constant.ADMIN_ROLE && payload.Username != req.Username {
		return nil, errors.New("unauthorized")
	}

	// Update user info
	user := &pb.UpdateUserInfoRequest{
		Username: req.Username,
		FullName: req.FullName,
		Phone:    req.Phone,
	}

	if err := server.store.Table("users").Where("username = ?", req.Username).Updates(user).Error; err != nil {
		return nil, err
	}

	// Delete user info in redis
	redisKey := fmt.Sprintf("%s:%s", constant.USER_REDIS, req.Username)
	if err := server.cache.Clear(ctx, redisKey); err != nil {
		log.Print(err)
	}

	// Delete all user info in redis
	redisKey = fmt.Sprintf("%s:all", constant.USER_REDIS)
	if err := server.cache.Clear(ctx, redisKey); err != nil {
		log.Print(err)
	}

	return &pb.UpdateUserInfoResponse{
		Success: true,
	}, nil
}
