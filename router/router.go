package router

import (
	"github.com/gin-gonic/gin"
	"github.com/travor-backend/db"
	"github.com/travor-backend/middleware"
	"github.com/travor-backend/util"
)

type Server struct {
	config util.Config
	router *gin.Engine
}

func NewServer(config util.Config) (*Server, error) {
	server := &Server{config: config}

	_, err := db.GetInstance(config.DBSource)
	if err != nil {
		return nil, err
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(middleware.CorsMiddleware())

	api := router.Group("/api/v1")

	{
		destinationRoutes(api)
		packagesRoutes(api)
		typesRoutes(api)
	}

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
