package router

import (
	"github.com/gin-gonic/gin"
	"github.com/travor-backend/controller"
	"github.com/travor-backend/db"
)

func packagesRoutes(router *gin.RouterGroup) {
	packages := router.Group("/packages")

	{
		packages.GET("/", controller.GetPackages(db.DB))
		packages.GET("/:id", controller.GetPackageById(db.DB))
		packages.PUT("/", controller.CreatePackage(db.DB))
		packages.POST("/:id", controller.UpdatePackage(db.DB))
		packages.DELETE("/:id", controller.DeletePackage(db.DB))
	}
}
