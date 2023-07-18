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

type PackageServer struct {
	cache  util.RedisUtil
	config util.Config
	store  *gorm.DB
	pb.UnimplementedPackageServiceServer
}

func NewPackageServer(config util.Config, cache util.RedisUtil) (*PackageServer, error) {
	pkgDb, err := db.Connect(config.GetDBSource(constant.DESTINATION_DB))
	if err != nil {
		return nil, err
	}

	return &PackageServer{store: pkgDb, config: config, cache: cache}, nil
}

func (server *PackageServer) IsValidPackage(ctx context.Context, req *pb.PackageID) (*pb.IsValidPackageResponse, error) {
	var packageModel model.Package

	if err := server.store.First(&packageModel, req.Id).Error; err != nil {
		return &pb.IsValidPackageResponse{
			Valid: false,
		}, nil
	}

	return &pb.IsValidPackageResponse{
		Valid: true,
	}, nil
}

func (server *PackageServer) GetPackages(ctx context.Context, req *pb.Pagination) (*pb.GetPackagesResponse, error) {
	// Create a Redis key based on the request parameters
	redisKey := fmt.Sprintf("%s:%d:%d", constant.PACKAGE_REDIS, req.PageSize, req.Page)

	// Get data from cache
	cachedResponse, err := server.cache.Get(ctx, redisKey)
	if err == nil {
		// If the response is found in the cache, deserialize it and return
		var res *pb.GetPackagesResponse
		err = json.Unmarshal([]byte(cachedResponse), &res)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	// If not found in cache, get from db
	var packages []*model.Package

	if err := server.store.Limit(int(req.PageSize)).Offset(int(req.PageSize) * (int(req.Page) - 1)).Find(&packages).Error; err != nil {
		return nil, err
	}

	makePackages := make([]model.Package, len(packages))
	for i, p := range packages {
		makePackages[i] = *p
	}

	convertedPackages := convertPackages(makePackages)

	var total int64
	if err := server.store.Model(&model.Package{}).Count(&total).Error; err != nil {
		return nil, err
	}

	res := &pb.GetPackagesResponse{
		Packages: convertedPackages,
		Pagination: &pb.PaginationRes{
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}

	// Serialize the response and store it in Redis cache for future use
	serializedResponse, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	err = server.cache.Set(ctx, redisKey, serializedResponse, time.Hour)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (server *PackageServer) GetPackage(ctx context.Context, req *pb.PackageID) (*pb.Package, error) {
	redisKey := fmt.Sprintf("%s#%d", constant.PACKAGE_REDIS, req.Id)

	// Get data from cache
	cachedResponse, err := server.cache.Get(ctx, redisKey)
	if err == nil {
		// If the response is found in the cache, deserialize it and return
		var packageModel model.Package
		err = json.Unmarshal([]byte(cachedResponse), &packageModel)
		if err != nil {
			return nil, err
		}
		return convertPackage(packageModel), nil
	}

	// If not found in cache, get from db
	var packageModel model.Package

	if err := server.store.First(&packageModel, req.Id).Error; err != nil {
		return nil, err
	}

	// Serialize the response and store it in Redis cache for future use
	serializedResponse, err := json.Marshal(packageModel)
	if err != nil {
		log.Println(err)
	}
	err = server.cache.Set(ctx, redisKey, serializedResponse, time.Hour)
	if err != nil {
		log.Println(err)
	}

	return convertPackage(packageModel), nil
}

func (server *PackageServer) CreatePackage(ctx context.Context, req *pb.CreatePackageReq) (*pb.Package, error) {
	// Add admin interceptor
	c, err := interceptor.AdminInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	packageModel := model.Package{
		Name:         req.Name,
		Details:      req.Details,
		Price:        req.Price,
		ImgURL:       req.ImgUrl,
		Duration:     req.Duration,
		NumberPeople: req.NumberPeople,
		DesID:        req.DesId,
	}

	if err := server.store.Create(&packageModel).Error; err != nil {
		return nil, err
	}

	for _, t := range req.Types {
		if err := server.store.Create(&model.PackageType{
			PID: packageModel.ID,
			TID: t,
		}).Error; err != nil {
			return nil, err
		}
	}

	convertedPackage := convertPackage(packageModel)

	// Clear and update cache
	if err = server.cache.ClearWithPrefix(ctx, constant.PACKAGE_REDIS+":"); err != nil {
		log.Println(err)
	}

	jsonPkg, err := json.Marshal(convertedPackage)
	if err != nil {
		log.Println(err)
	}

	if err = server.cache.Set(ctx, fmt.Sprintf("%s#%d", constant.PACKAGE_REDIS, convertedPackage.Id), jsonPkg, time.Hour); err != nil {
		log.Println(err)
	}
	return convertedPackage, nil
}

func (server *PackageServer) UpdatePackage(ctx context.Context, req *pb.UpdatePackageReq) (*pb.Package, error) {
	// Add admin interceptor
	c, err := interceptor.AdminInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	packageModel := model.Package{
		Name:         req.Name,
		Details:      req.Details,
		Price:        req.Price,
		ImgURL:       req.ImgUrl,
		Duration:     req.Duration,
		NumberPeople: req.NumberPeople,
		DesID:        req.DesId,
	}

	if err := server.store.Model(&packageModel).Where("id = ?", req.Id).Updates(&packageModel).Error; err != nil {
		return nil, err
	}

	if err := server.store.Where("p_id = ?", req.Id).Delete(&model.PackageType{}).Error; err != nil {
		return nil, err
	}

	for _, t := range req.Types {
		if err := server.store.Create(&model.PackageType{
			PID: req.Id,
			TID: t,
		}).Error; err != nil {
			return nil, err
		}
	}

	convertedPackage := convertPackage(packageModel)

	// Clear and update cache
	if err = server.cache.ClearWithPrefix(ctx, constant.PACKAGE_REDIS+":"); err != nil {
		log.Println(err)
	}

	jsonPkg, err := json.Marshal(convertedPackage)
	if err != nil {
		log.Println(err)
	}

	if err = server.cache.Set(ctx, fmt.Sprintf("%s#%d", constant.PACKAGE_REDIS, req.Id), jsonPkg, time.Hour); err != nil {
		log.Println(err)
	}

	return convertedPackage, nil
}

func (server *PackageServer) DeletePackage(ctx context.Context, req *pb.PackageID) (*pb.DeletePackageResponse, error) {
	// Add admin interceptor
	c, err := interceptor.AdminInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	var packageModel model.Package

	if err := server.store.Delete(model.PackageType{}, "p_id = ?", req.Id).Error; err != nil {
		return &pb.DeletePackageResponse{
			Success: false,
		}, nil
	}

	if err := server.store.Delete(&packageModel, req.Id).Error; err != nil {
		return &pb.DeletePackageResponse{
			Success: false,
		}, nil
	}

	if err = server.cache.ClearWithPrefix(ctx, constant.PACKAGE_REDIS+":"); err != nil {
		log.Println(err)
	}

	if err = server.cache.Clear(ctx, fmt.Sprintf("%s#%d", constant.PACKAGE_REDIS, req.Id)); err != nil {
		log.Println(err)
	}

	return &pb.DeletePackageResponse{
		Success: true,
	}, nil
}
