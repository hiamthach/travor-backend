package router

import (
	"github.com/gin-gonic/gin"
	"github.com/travor-backend/config"
)

type Server struct {
	config config.Config
	router *gin.Engine
}

func NewServer(config config.Config) (*Server, error) {
	server := &Server{config: config}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	api := router.Group("/api/v1")

	{
		destinationRoutes(api)
	}

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
