package gapi

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/travor-backend/pb"
	"github.com/travor-backend/util"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/gorm"
)

func RunGRPCServer(config util.Config, store *gorm.DB, cache util.RedisUtil) {
	grpcServer := grpc.NewServer()

	// Register auth server
	authServer, err := NewAuthServer(config)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}

	pb.RegisterAuthServiceServer(grpcServer, authServer)

	// Register destination server
	destinationServer, err := NewDestinationServer(config, cache)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}

	pb.RegisterDestinationServiceServer(grpcServer, destinationServer)

	// Register type server
	typeServer, err := NewTypeServer(config, cache)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}

	pb.RegisterTypeServiceServer(grpcServer, typeServer)

	// Register package server
	packageServer, err := NewPackageServer(config, cache)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}

	pb.RegisterPackageServiceServer(grpcServer, packageServer)

	// Start server
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Can not start server: ", err)
	}

	log.Println("Starting gRPC server on", config.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Can not start server: ", err)
	}
}

func RunGatewayServer(config util.Config, store *gorm.DB, cache util.RedisUtil) {
	// initialize grpc server
	authServer, err := NewAuthServer(config)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}
	destinationServer, err := NewDestinationServer(config, cache)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}

	typeServer, err := NewTypeServer(config, cache)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}

	packageServer, err := NewPackageServer(config, cache)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	// initialize gateway server
	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// register gateway server
	if err = pb.RegisterAuthServiceHandlerServer(ctx, grpcMux, authServer); err != nil {
		log.Fatal("Can not register gateway server: ", err)
	}

	if err = pb.RegisterDestinationServiceHandlerServer(ctx, grpcMux, destinationServer); err != nil {
		log.Fatal("Can not register gateway server: ", err)
	}

	if err = pb.RegisterTypeServiceHandlerServer(ctx, grpcMux, typeServer); err != nil {
		log.Fatal("Can not register gateway server: ", err)
	}

	if err = pb.RegisterPackageServiceHandlerServer(ctx, grpcMux, packageServer); err != nil {
		log.Fatal("Can not register gateway server: ", err)
	}

	// initialize http server
	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", grpcMux))

	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatal("Can not start server: ", err)
	}

	log.Println("Starting gateway server on", config.ServerAddress)

	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("Can not start server: ", err)
	}
}
