package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/travor-backend/db"
	_ "github.com/travor-backend/docs"
	"github.com/travor-backend/gapi"
	"github.com/travor-backend/pb"
	"github.com/travor-backend/util"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/gorm"
)

// @title           Travor Backend API
// @version         1.0

// @BasePath  /api/v1
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not load config: ", err)
	}

	// Run Gin Server
	// server, err := router.NewServer(config)
	// if err != nil {
	// 	log.Fatal("Can not create server: ", err)
	// }

	// err = server.Start(config.ServerAddress)
	// if err != nil {
	// 	log.Fatal("Can not start server: ", err)
	// }

	_, err = db.GetInstance(config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to database: ", err)
	}

	go runGatewayServer(config, db.DB)
	runGRPCServer(config, db.DB)
}

func runGRPCServer(config util.Config, store *gorm.DB) {
	server, err := gapi.NewGRPCServer(store, config)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, server)

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

func runGatewayServer(config util.Config, store *gorm.DB) {
	server, err := gapi.NewGRPCServer(store, config)
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

	err = pb.RegisterAuthServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("Can not register gateway server: ", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

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
