package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/travor-backend/db"
	"github.com/travor-backend/middleware"
	"github.com/travor-backend/util"
)

type Server struct {
	config util.Config
	router *gin.Engine
}

var Config util.Config

func NewServer(config util.Config) (*Server, error) {
	var err error

	Config = config
	server := &Server{config: config}

	_, err = db.GetInstance(config.DBSource)
	if err != nil {
		return nil, err
	}

	err = db.InitializeFirebase()
	if err != nil {
		return nil, err
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.New()
	router.Use(middleware.CorsMiddleware())

	api := router.Group("/api/v1")

	{
		packageRoutes(api)
		typeRoutes(api)
		galleryRoutes(api)
		userRoutes(api)
		tripRoutes(api)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
