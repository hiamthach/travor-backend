package gapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/travor-backend/constant"
	"github.com/travor-backend/db"
	"github.com/travor-backend/interceptor"
	"github.com/travor-backend/pb"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

type DestinationServer struct {
	cache  util.RedisUtil
	config util.Config
	store  *gorm.DB
	pb.UnimplementedDestinationServiceServer
}

func NewDestinationServer(config util.Config, cache util.RedisUtil) (*DestinationServer, error) {
	desDb, err := db.Connect(config.GetDBSource(constant.DESTINATION_DB))
	if err != nil {
		return nil, err
	}

	return &DestinationServer{store: desDb, config: config, cache: cache}, nil
}

func (server *DestinationServer) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	// Get from redis
	redisKey := fmt.Sprintf("%s:all", constant.DESTINATION_REDIS)

	cachedResponse, err := server.cache.Get(ctx, redisKey)
	if err == nil {
		// If the response is found in the cache, deserialize it and return
		var response *pb.GetAllResponse
		err = json.Unmarshal([]byte(cachedResponse), &response)
		if err != nil {
			return nil, err
		}
		return response, err
	}

	var destinations []*pb.DestinationStats

	if err := server.store.Table("destinations").Select("id", "name").Find(&destinations).Error; err != nil {
		return nil, err
	}

	res := &pb.GetAllResponse{
		Destinations: destinations,
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

func (server *DestinationServer) GetDestinations(ctx context.Context, req *pb.GetDestinationsRequest) (*pb.GetDestinationsResponse, error) {
	// Create a Redis key based on the request parameters
	var redisKey string
	if req.Keyword == "" {
		redisKey = fmt.Sprintf("%s:%d:%d", constant.DESTINATION_REDIS, req.PageSize, req.Page)
	} else {
		redisKey = fmt.Sprintf("%s:%d:%d:%s", constant.DESTINATION_REDIS, req.PageSize, req.Page, req.Keyword)
	}

	// Get data from cache
	cachedResponse, err := server.cache.Get(ctx, redisKey)
	if err == nil {
		// If the response is found in the cache, deserialize it and return
		var response *pb.GetDestinationsResponse
		err = json.Unmarshal([]byte(cachedResponse), &response)
		if err != nil {
			return nil, err
		}
		return response, err
	}

	// If not found in cache, get from db
	var destinations []*pb.Destination

	// If keyword is empty, get all destinations, else get destinations by keyword ignore case
	if err := server.store.
		Limit(int(req.PageSize)).
		Offset(int(req.PageSize)*(int(req.Page)-1)).
		Where("LOWER(name) LIKE ?", "%"+strings.ToLower(req.Keyword)+"%").
		Find(&destinations).Error; err != nil {
		return nil, err
	}

	var total int64
	if err := server.store.Model(&pb.Destination{}).Where("LOWER(name) LIKE ?", "%"+strings.ToLower(req.Keyword)+"%").Count(&total).Error; err != nil {
		return nil, err
	}

	res := &pb.GetDestinationsResponse{
		Destinations: destinations,
		Pagination: &pb.PaginationRes{
			PageSize: req.PageSize,
			Page:     req.Page,
			Total:    total,
		},
	}

	// Serialize the response and store it in Redis cache for future use
	serializedResponse, err := json.Marshal(res)
	if err != nil {
		log.Print(err)
	}
	err = server.cache.Set(ctx, redisKey, serializedResponse, time.Hour)
	if err != nil {
		log.Print(err)
	}

	return res, nil
}

func (server *DestinationServer) GetDestinationById(ctx context.Context, req *pb.GetDestinationByIdRequest) (*pb.Destination, error) {
	// Create a Redis key based on the request parameters
	redisKey := fmt.Sprintf("%s#%d", constant.DESTINATION_REDIS, req.Id)

	// Get data from cache
	cachedResponse, err := server.cache.Get(ctx, redisKey)
	if err == nil {
		// If the response is found in the cache, deserialize it and return
		var destination *pb.Destination
		err = json.Unmarshal([]byte(cachedResponse), &destination)
		if err != nil {
			return nil, err
		}
		return destination, nil
	}

	// If not found in cache, get from db
	var destination *pb.Destination

	if err := server.store.First(&destination, req.Id).Error; err != nil {
		return nil, err
	}

	// Serialize the response and store it in Redis cache for future use
	serializedResponse, err := json.Marshal(destination)
	if err != nil {
		log.Print(err)
	}
	err = server.cache.Set(ctx, redisKey, serializedResponse, time.Hour)
	if err != nil {
		log.Print(err)
	}

	return destination, nil
}

func (server *DestinationServer) CreateDestination(ctx context.Context, req *pb.CreateDestinationRequest) (*pb.CreateDestinationResponse, error) {
	// Add admin interceptor
	c, err := interceptor.AdminInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	destination := &pb.Destination{
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		Price:       req.Price,
		Language:    req.Language,
		VisaRequire: req.VisaRequire,
		Country:     req.Country,
		Currency:    req.Currency,
		Area:        req.Area,
	}

	if err := server.store.Create(&destination).Error; err != nil {
		return nil, err
	}

	// Clear cache
	if err = server.cache.ClearWithPrefix(ctx, constant.DESTINATION_REDIS+":"); err != nil {
		log.Println("Error clearing cache: ", err)
	}
	// Set cache data for new destination
	jsonDes, err := json.Marshal(destination)
	if err != nil {
		log.Print(err)
	}

	if err = server.cache.Set(ctx, fmt.Sprintf("%s#%d", constant.DESTINATION_REDIS, destination.Id), jsonDes, time.Hour); err != nil {
		log.Println("Set cache error: ", err)
	}

	return &pb.CreateDestinationResponse{
		Destination: destination,
	}, nil
}

func (server *DestinationServer) UpdateDestination(ctx context.Context, req *pb.UpdateDestinationRequest) (*pb.UpdateDestinationResponse, error) {
	// Add admin interceptor
	c, err := interceptor.AdminInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	// If the req not have enough field it not update field not in req
	destination := &pb.Destination{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		Price:       req.Price,
		Language:    req.Language,
		VisaRequire: req.VisaRequire,
		Country:     req.Country,
		Currency:    req.Currency,
		Area:        req.Area,
	}

	if err := server.store.Model(&pb.Destination{}).Where("id = ?", req.Id).Updates(destination).Error; err != nil {
		return nil, err
	}

	// Clear cache
	if err = server.cache.ClearWithPrefix(ctx, constant.DESTINATION_REDIS+":"); err != nil {
		log.Println("Error clearing cache: ", err)
	}
	// Set cache data for new destination
	jsonDes, err := json.Marshal(destination)
	if err != nil {
		log.Print(err)
	}

	if err = server.cache.Set(ctx, fmt.Sprintf("%s#%d", constant.DESTINATION_REDIS, destination.Id), jsonDes, time.Hour); err != nil {
		log.Println("Set cache error: ", err)
	}

	return &pb.UpdateDestinationResponse{
		Destination: destination,
	}, nil
}

func (server *DestinationServer) DeleteDestination(ctx context.Context, req *pb.DeleteDestinationRequest) (*pb.DeleteDestinationResponse, error) {
	// Add admin interceptor
	c, err := interceptor.AdminInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	destination := &pb.Destination{
		Id: req.Id,
	}

	if err := server.store.Delete(&destination).Error; err != nil {
		return nil, err
	}

	// Clear cache
	if err = server.cache.ClearWithPrefix(ctx, constant.DESTINATION_REDIS+":"); err != nil {
		log.Println("Error clearing cache: ", err)
	}

	if err = server.cache.Clear(ctx, fmt.Sprintf("%s#%d", constant.DESTINATION_REDIS, destination.Id)); err != nil {
		log.Println("Error clearing cache: ", err)
	}

	return &pb.DeleteDestinationResponse{
		Success: true,
	}, nil
}
