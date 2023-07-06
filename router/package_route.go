package router

import (
	"github.com/gin-gonic/gin"
	"github.com/travor-backend/controller"
	"github.com/travor-backend/db"
	"github.com/travor-backend/middleware"
)

func packageRoutes(router *gin.RouterGroup) {
	packages := router.Group("/packages")

	{
		packages.GET("/", controller.GetPackages(db.DB))
		packages.GET("/:id", controller.GetPackageById(db.DB))
		// middleware with auth and admin
		packages.Use(middleware.AuthMiddleware())
		packages.Use(middleware.AdminMiddleware(db.DB))
		packages.POST("/", controller.CreatePackage(db.DB))
		packages.PUT("/:id", controller.UpdatePackage(db.DB))
		packages.DELETE("/:id", controller.DeletePackage(db.DB))
	}
}
