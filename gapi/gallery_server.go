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
	"github.com/travor-backend/pb"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

type GalleryServer struct {
	cache  util.RedisUtil
	config util.Config
	store  *gorm.DB
	pb.UnimplementedGalleryServiceServer
}

func NewGalleryServer(config util.Config, cache util.RedisUtil) (*GalleryServer, error) {
	galDb, err := db.Connect(config.GetDBSource(constant.GALLERY_DB))
	if err != nil {
		return nil, err
	}

	return &GalleryServer{store: galDb, config: config, cache: cache}, nil
}

func (server *GalleryServer) AddImage(ctx context.Context, req *pb.AddImageRequest) (*pb.AddImageResponse, error) {
	// Add admin interceptor
	c, err := interceptor.AdminInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	gallery := &pb.Image{
		DesId: req.DesId,
		Url:   req.Url,
	}

	if err := server.store.Table("galleries").Create(&gallery).Error; err != nil {
		return nil, err
	}

	// Clear cache
	redisKey := fmt.Sprintf("%s:%d", constant.GALLERY_REDIS, req.DesId)
	if err := server.cache.Clear(ctx, redisKey); err != nil {
		log.Println(err)
	}

	return &pb.AddImageResponse{
		Image: gallery,
	}, nil
}

func (server *GalleryServer) GetImages(ctx context.Context, req *pb.GetImagesRequest) (*pb.GetImagesResponse, error) {
	redisKey := fmt.Sprintf("%s:%d", constant.GALLERY_REDIS, req.DesId)

	// Get data from cache
	cachedResponse, err := server.cache.Get(ctx, redisKey)
	if err == nil {
		// If the response is found in the cache, deserialize it and return
		var images []*pb.Image
		err = json.Unmarshal([]byte(cachedResponse), &images)
		if err != nil {
			return nil, err
		}
		return &pb.GetImagesResponse{
			Images: images,
		}, nil
	}

	var images []*pb.Image

	if err := server.store.Table("galleries").Where("des_id = ?", req.DesId).Find(&images).Error; err != nil {
		return nil, err
	}

	// Serialize the response
	serializedResponse, err := json.Marshal(images)
	if err != nil {
		log.Println(err)
	}

	// Set the serialized response in the cache, expire in 1 hour
	if err := server.cache.Set(ctx, redisKey, serializedResponse, time.Hour); err != nil {
		log.Println(err)
	}

	return &pb.GetImagesResponse{
		Images: images,
	}, nil
}

func (server *GalleryServer) DeleteImage(ctx context.Context, req *pb.DeleteImageRequest) (*pb.DeleteImageResponse, error) {
	// Add admin interceptor
	c, err := interceptor.AdminInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	if err := server.store.Table("galleries").Where("id = ?", req.Id).Delete(&pb.Image{}).Error; err != nil {
		return nil, err
	}

	// Clear cache
	if err := server.cache.ClearWithPrefix(ctx, constant.GALLERY_REDIS+":"); err != nil {
		log.Println(err)
	}

	return &pb.DeleteImageResponse{
		Success: true,
	}, nil
}
