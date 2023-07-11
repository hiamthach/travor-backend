package main

import (
	"log"

	"github.com/travor-backend/db"
	"github.com/travor-backend/gapi"
	"github.com/travor-backend/util"
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

	// Run gRPC Server
	go gapi.RunGatewayServer(config, db.DB, *redisUtil)
	gapi.RunGRPCServer(config, db.DB, *redisUtil)
}
