package router

import (
	"github.com/gin-gonic/gin"
	"github.com/travor-backend/controller"
	"github.com/travor-backend/db"
	"github.com/travor-backend/middleware"
)

func typeRoutes(router *gin.RouterGroup) {
	packages := router.Group("/types")

	{
		packages.GET("/", controller.GetTypes(db.DB))
		packages.GET("/:id", controller.GetTypeById(db.DB))
		// middleware with auth and admin
		packages.Use(middleware.AuthMiddleware(TokenMaker))
		packages.Use(middleware.AdminMiddleware(db.DB))
		packages.POST("/", controller.CreateType(db.DB))
		packages.PUT("/:id", controller.UpdateType(db.DB))
		packages.DELETE("/:id", controller.DeleteType(db.DB))
	}
}
