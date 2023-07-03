package gapi

import (
	"context"

	"github.com/travor-backend/pb"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

type GRPCServer struct {
	config util.Config
	store  *gorm.DB
	pb.UnimplementedAuthServiceServer
}

func NewGRPCServer(store *gorm.DB, config util.Config) (*GRPCServer, error) {
	return &GRPCServer{store: store, config: config}, nil
}

func (server *GRPCServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		Token: "token",
	}, nil
}
