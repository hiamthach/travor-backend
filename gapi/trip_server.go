package gapi

import (
	"context"
	"errors"
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
		User:      payload.Username,
		PId:       req.PId,
		Total:     req.Total,
		Notes:     req.Notes,
		StartDate: startDate,
	}

	// Insert Trip object into database
	if err := server.store.Table("trips").Create(&trip).Error; err != nil {
		return nil, err
	}

	return &pb.Trip{
		Id:        trip.Id,
		User:      trip.User,
		PId:       trip.PId,
		Total:     trip.Total,
		Notes:     trip.Notes,
		StartDate: timestamppb.New(trip.StartDate),
	}, nil
}

func (server *TripServer) GetTrip(ctx context.Context, req *pb.GetTripRequest) (*pb.Trip, error) {
	return nil, nil
}

func (server *TripServer) GetTrips(ctx context.Context, req *pb.GetTripsRequest) (*pb.GetTripsResponse, error) {
	return nil, nil
}

func (server *TripServer) UpdateTrip(ctx context.Context, req *pb.UpdateTripRequest) (*pb.Trip, error) {
	return nil, nil
}

func (server *TripServer) DeleteTrip(ctx context.Context, req *pb.DeleteTripRequest) (*pb.DeleteTripResponse, error) {
	return nil, nil
}
