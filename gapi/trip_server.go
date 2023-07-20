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
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type TripServer struct {
	cache  util.RedisUtil
	config util.Config
	store  *gorm.DB
	pb.UnimplementedTripServiceServer
	clientConn *grpc.ClientConn
}

func NewTripServer(config util.Config, cache util.RedisUtil, conn *grpc.ClientConn) (*TripServer, error) {
	desDb, err := db.Connect(config.GetDBSource(constant.TRIP_DB))
	if err != nil {
		return nil, err
	}

	return &TripServer{store: desDb, config: config, cache: cache, clientConn: conn}, nil
}

func (server *TripServer) CreateTrip(ctx context.Context, req *pb.CreateTripRequest) (*pb.Trip, error) {
	// Authenticate the incoming request
	c, err := interceptor.AuthInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	// Check if the package exists
	packageClient := pb.NewPackageServiceClient(server.clientConn)

	valid, err := packageClient.IsValidPackage(ctx, &pb.PackageID{Id: req.PId})
	if err != nil {
		return nil, err
	}

	if !valid.Valid {
		return nil, errors.New("invalid package")
	}

	// Create Trip object
	startDate, err := time.Parse(time.RFC3339, req.StartDate)
	if err != nil {
		return nil, err
	}

	trip := &model.Trip{
		Username:  payload.Username,
		PId:       req.PId,
		Total:     req.Total,
		Notes:     req.Notes,
		StartDate: startDate,
	}

	// Insert Trip object into database
	if err := server.store.Table("trips").Create(&trip).Error; err != nil {
		return nil, err
	}

	// Clear cache
	redisKey := fmt.Sprintf("%s:%s", constant.TRIP_REDIS, payload.Username)
	if err := server.cache.ClearWithPrefix(ctx, redisKey); err != nil {
		log.Printf("Failed to delete cache: %v\n", err)
	}

	return &pb.Trip{
		Id:        trip.Id,
		Username:  trip.Username,
		PId:       trip.PId,
		Total:     trip.Total,
		Notes:     trip.Notes,
		StartDate: timestamppb.New(trip.StartDate),
	}, nil
}

func (server *TripServer) GetTrip(ctx context.Context, req *pb.GetTripRequest) (*pb.Trip, error) {
	c, err := interceptor.AuthInterceptor(ctx)
	if err != nil {
		return nil, err
	}
	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	// Get the trip from the database
	var trip *model.Trip

	if err := server.store.Table("trips").Where("id = ?", req.Id).First(&trip).Error; err != nil {
		return nil, err
	}
	// Convert the trip to a protocol buffer object
	pbTrip := &pb.Trip{
		Id:        trip.Id,
		Username:  trip.Username,
		PId:       trip.PId,
		Total:     trip.Total,
		Notes:     trip.Notes,
		StartDate: timestamppb.New(trip.StartDate),
	}

	return pbTrip, nil
}

func (server *TripServer) GetTrips(ctx context.Context, req *pb.GetTripsRequest) (*pb.GetTripsResponse, error) {
	// Authenticate the incoming request
	c, err := interceptor.AuthInterceptor(ctx)
	if err != nil {
		return nil, err
	}

	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	// Redis cache
	redisKey := fmt.Sprintf("%s:%s", constant.TRIP_REDIS, payload.Username)

	// Get data from cache
	cachedResponse, err := server.cache.Get(ctx, redisKey)
	if err == nil {
		// If the response is found in the cache, deserialize it and return
		var response *pb.GetTripsResponse
		err = json.Unmarshal([]byte(cachedResponse), &response)
		if err != nil {
			return nil, err
		}
		return response, err
	}

	// Get all trips for the user
	var trips []*model.Trip

	if err := server.store.Table("trips").Where("username = ?", payload.Username).Find(&trips).Error; err != nil {
		return nil, err
	}
	// Convert trips to protocol buffer objects
	var pbTrips []*pb.Trip
	for _, trip := range trips {
		pbTrips = append(pbTrips, &pb.Trip{
			Id:        trip.Id,
			Username:  trip.Username,
			PId:       trip.PId,
			Total:     trip.Total,
			Notes:     trip.Notes,
			StartDate: timestamppb.New(trip.StartDate),
		})
	}

	res := &pb.GetTripsResponse{Trips: pbTrips}
	// Serialize the response
	serializedResponse, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	// Save the response to cache
	if err := server.cache.Set(ctx, redisKey, string(serializedResponse), time.Hour); err != nil {
		log.Printf("Failed to cache trips: %v\n", err)
	}

	return res, nil
}

func (server *TripServer) UpdateTrip(ctx context.Context, req *pb.UpdateTripRequest) (*pb.Trip, error) {
	// Authenticate the incoming request
	c, err := interceptor.AuthInterceptor(ctx)
	if err != nil {
		return nil, err
	}
	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	// Get the trip from the database
	// var startDate  *time.Time
	var trip *model.Trip

	if req.StartDate != "" {
		startDate, err := time.Parse(time.RFC3339, req.StartDate)
		if err != nil {
			return nil, err
		}

		trip = &model.Trip{
			Id:        req.Id,
			Username:  payload.Username,
			Total:     req.Total,
			Notes:     req.Notes,
			StartDate: startDate,
		}
	} else {
		trip = &model.Trip{
			Id:        req.Id,
			Username:  payload.Username,
			Total:     req.Total,
			Notes:     req.Notes,
			StartDate: time.Time{},
		}
	}

	// Update the trip in the database
	if err := server.store.Table("trips").Where("id = ?", req.Id).Updates(&trip).Error; err != nil {
		return nil, err
	}

	// Clear cache
	redisKey := fmt.Sprintf("%s:%s", constant.TRIP_REDIS, payload.Username)
	if err := server.cache.ClearWithPrefix(ctx, redisKey); err != nil {
		log.Printf("Failed to delete cache: %v\n", err)
	}

	return convertTrip(*trip), nil
}

func (server *TripServer) DeleteTrip(ctx context.Context, req *pb.DeleteTripRequest) (*pb.DeleteTripResponse, error) {
	// Authenticate the incoming request
	c, err := interceptor.AdminInterceptor(ctx)
	if err != nil {
		return nil, err
	}
	payload := c.Value(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
	if payload == nil {
		return nil, errors.New("unauthorized")
	}

	// Delete the trip from the database

	var trip *model.Trip
	if err := server.store.Table("trips").Where("id = ?", req.Id).Delete(&trip).Error; err != nil {
		return &pb.DeleteTripResponse{Success: false}, err
	}

	// Clear cache
	redisKey := fmt.Sprintf("%s:", constant.TRIP_REDIS)
	if err := server.cache.Clear(ctx, redisKey); err != nil {
		log.Printf("Failed to delete cache: %v\n", err)
	}

	return &pb.DeleteTripResponse{Success: true}, nil
}
