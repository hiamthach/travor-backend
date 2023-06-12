package router

import (
	"github.com/gin-gonic/gin"
	"github.com/travor-backend/controller"
	"github.com/travor-backend/db"
)

func typesRoutes(router *gin.RouterGroup) {
	packages := router.Group("/types")

	{
		packages.GET("/", controller.GetTypes(db.DB))
		packages.GET("/:id", controller.GetTypeById(db.DB))
		packages.PUT("/", controller.CreateType(db.DB))
		packages.POST("/:id", controller.UpdateType(db.DB))
		packages.DELETE("/:id", controller.DeleteType(db.DB))
	}
}
