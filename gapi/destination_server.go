package gapi

import (
	"context"

	"github.com/travor-backend/pb"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

type DestinationServer struct {
	config util.Config
	store  *gorm.DB
	pb.UnimplementedDestinationServiceServer
}

func NewDestinationServer(store *gorm.DB, config util.Config) (*DestinationServer, error) {
	return &DestinationServer{store: store, config: config}, nil
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
