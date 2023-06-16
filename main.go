package main

import (
	"log"

	_ "github.com/travor-backend/docs"
	"github.com/travor-backend/router"
	"github.com/travor-backend/util"
)

// @title           Travor Backend API
// @version         1.0

// @host      localhost:8088
// @BasePath  /api/v1
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not load config: ", err)
	}

	server, err := router.NewServer(config)
	if err != nil {
		log.Fatal("Can not create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can not start server: ", err)
	}
}
