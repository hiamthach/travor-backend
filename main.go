package main

import (
	"log"

	"github.com/travor-backend/db"
	_ "github.com/travor-backend/docs"
	"github.com/travor-backend/gapi"
	"github.com/travor-backend/util"
)

// @title           Travor Backend API
// @version         1.0

// @BasePath  /api/v1
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not load config: ", err)
	}
	_, err = db.GetInstance(config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to database: ", err)
	}

	// Run Gin Server
	// go func() {
	// 	server, err := router.NewServer(config)
	// 	if err != nil {
	// 		log.Fatal("Can not create server: ", err)
	// 	}

	// 	err = server.Start(config.ServerAddress)
	// 	if err != nil {
	// 		log.Fatal("Can not start server: ", err)
	// 	}
	// }()

	// Run gRPC Server
	go gapi.RunGatewayServer(config, db.DB)
	gapi.RunGRPCServer(config, db.DB)
}
