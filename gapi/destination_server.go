package gapi

import (
	"context"
	"errors"

	"github.com/travor-backend/constant"
	"github.com/travor-backend/db"
	"github.com/travor-backend/interceptor"
	"github.com/travor-backend/pb"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

type DestinationServer struct {
	config util.Config
	store  *gorm.DB
	pb.UnimplementedDestinationServiceServer
}

func NewDestinationServer(config util.Config) (*DestinationServer, error) {
	desDb, err := db.Connect(config.GetDBSource(constant.DESTINATION_DB))
	if err != nil {
		return nil, err
	}
	return &DestinationServer{store: desDb, config: config}, nil
}

func (server *DestinationServer) GetDestinations(ctx context.Context, req *pb.GetDestinationsRequest) (*pb.GetDestinationsResponse, error) {
	var destinations []*pb.Destination

	if err := server.store.Limit(int(req.PageSize)).Offset(int(req.PageSize) * (int(req.Page) - 1)).Find(&destinations).Error; err != nil {
		return nil, err
	}

	return &pb.GetDestinationsResponse{
		Destinations: destinations,
	}, nil
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

	return &pb.CreateDestinationResponse{
		Destination: destination,
	}, nil
}
