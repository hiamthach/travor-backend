package router

import (
	"github.com/gin-gonic/gin"
	"github.com/travor-backend/controller"
	"github.com/travor-backend/db"
)

func tripRoutes(router *gin.RouterGroup) {
	trips := router.Group("/trips")

	{
		trips.GET("/", controller.GetTrips(db.DB))
		trips.GET("/:id", controller.GetTripById(db.DB))
		trips.POST("/", controller.CreateTrip(db.DB))
		trips.PUT("/:id", controller.UpdateTrip(db.DB))
		trips.DELETE("/:id", controller.DeleteTrip(db.DB))
	}
}
