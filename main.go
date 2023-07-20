package main

import (
	"log"

	"github.com/travor-backend/db"
	"github.com/travor-backend/gapi"
	"github.com/travor-backend/util"
	"google.golang.org/grpc"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not load config: ", err)
	}

	err = db.InitializeFirebase()
	if err != nil {
		log.Fatal("Can not initialize firebase: ", err)
	}

	redisUtil, err := util.NewRedisUtil(config)
	if err != nil {
		log.Fatal("Can not connect to redis: ", err)
	}

	conn, err := grpc.Dial(config.GRPCServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("can not connect to grpc server: %w", err)
	}

	defer conn.Close()

	// Run gRPC Server
	go gapi.RunGRPCServer(config, db.DB, *redisUtil, conn)
	gapi.RunGatewayServer(config, db.DB, *redisUtil, conn)
}
