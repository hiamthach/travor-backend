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

func RunGRPCServer(config util.Config, store *gorm.DB) {
	authServer, err := NewAuthServer(store, config)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}

	destinationServer, err := NewDestinationServer(store, config)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, authServer)
	pb.RegisterDestinationServiceServer(grpcServer, destinationServer)

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

func RunGatewayServer(config util.Config, store *gorm.DB) {
	authServer, err := NewAuthServer(store, config)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}
	destinationServer, err := NewDestinationServer(store, config)
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

	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err = pb.RegisterAuthServiceHandlerServer(ctx, grpcMux, authServer); err != nil {
		log.Fatal("Can not register gateway server: ", err)
	}

	if err = pb.RegisterDestinationServiceHandlerServer(ctx, grpcMux, destinationServer); err != nil {
		log.Fatal("Can not register gateway server: ", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Can not start server: ", err)
	}

	log.Println("Starting gateway server on", ":8080")

	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("Can not start server: ", err)
	}
}
