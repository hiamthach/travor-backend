package router

import (
	"github.com/gin-gonic/gin"
	"github.com/travor-backend/controller"
	"github.com/travor-backend/db"
	"github.com/travor-backend/middleware"
)

func destinationRoutes(router *gin.RouterGroup) {
	destinations := router.Group("/destinations")

	{
		destinations.GET("/", controller.GetDestinations(db.DB))
		destinations.GET("/:id", controller.GetDestinationById(db.DB))
		// middleware with auth and admin
		destinations.Use(middleware.AuthMiddleware())
		destinations.Use(middleware.AdminMiddleware(db.DB))
		destinations.POST("/", controller.CreateDestination(db.DB))
		destinations.PUT("/:id", controller.UpdateDestination(db.DB))
		destinations.DELETE("/:id", controller.DeleteDestination(db.DB))
	}
}
