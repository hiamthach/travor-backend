package gapi

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/travor-backend/constant"
	"github.com/travor-backend/db"
	"github.com/travor-backend/pb"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

type TypeServer struct {
	cache  util.RedisUtil
	config util.Config
	store  *gorm.DB
	pb.UnimplementedTypeServiceServer
}

func NewTypeServer(config util.Config, cache util.RedisUtil) (*TypeServer, error) {
	desDb, err := db.Connect(config.GetDBSource(constant.DESTINATION_DB))
	if err != nil {
		return nil, err
	}

	return &TypeServer{store: desDb, config: config, cache: cache}, nil
}

func (server *TypeServer) GetTypes(ctx context.Context, req *pb.GetTypesRequest) (*pb.GetTypesResponse, error) {
	// Get data from caches
	redisKey := constant.TYPE_REDIS

	cachedResponse, err := server.cache.Get(ctx, redisKey)
	if err == nil {
		var types []*pb.Type
		err = json.Unmarshal([]byte(cachedResponse), &types)
		if err != nil {
			return nil, err
		}
		return &pb.GetTypesResponse{
			Types: types,
		}, nil
	}

	var types []*pb.Type

	if err := server.store.Table("types").Find(&types).Error; err != nil {
		return nil, err
	}

	// Serialize the response and store it in Redis cache for future use
	serializedResponse, err := json.Marshal(types)
	if err != nil {
		log.Print(err)
	}

	err = server.cache.Set(ctx, redisKey, serializedResponse, time.Hour)
	if err != nil {
		log.Print(err)
	}

	return &pb.GetTypesResponse{
		Types: types,
	}, nil
}

func (server *TypeServer) CreateType(ctx context.Context, req *pb.CreateTypeRequest) (*pb.CreateTypeResponse, error) {
	var typeObj pb.Type
	typeObj.Name = req.Name

	if err := server.store.Table("types").Create(&typeObj).Error; err != nil {
		return nil, err
	}

	// Clear cache
	if err := server.cache.ClearWithPrefix(ctx, constant.TYPE_REDIS); err != nil {
		log.Println("Error clearing cache: ", err)
	}

	return &pb.CreateTypeResponse{
		Success: true,
	}, nil
}

func (server *TypeServer) DeleteType(ctx context.Context, req *pb.DeleteTypeRequest) (*pb.DeleteTypeResponse, error) {
	if err := server.store.Table("types").Where("id = ?", req.Id).Delete(&pb.Type{}).Error; err != nil {
		return nil, err
	}

	// Clear cache
	if err := server.cache.ClearWithPrefix(ctx, constant.TYPE_REDIS); err != nil {
		log.Println("Error clearing cache: ", err)
	}

	return &pb.DeleteTypeResponse{
		Success: true,
	}, nil
}
