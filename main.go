package main

import (
	"log"

	"github.com/travor-backend/config"
	"github.com/travor-backend/router"
)

func main() {
	config, err := config.LoadConfig(".")
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
