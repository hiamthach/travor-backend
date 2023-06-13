package router

import (
	"github.com/gin-gonic/gin"
	"github.com/travor-backend/controller"
	"github.com/travor-backend/db"
)

func galleryRoutes(router *gin.RouterGroup) {
	galleries := router.Group("/galleries")

	{
		galleries.GET("/:des", controller.GetGalleriesByDesID(db.DB))
		galleries.POST("/upload/:des", controller.UploadToGallery(db.DB))
		// galleries.POST("/:id", controller.UpdateGallery(db.DB))
		galleries.DELETE("/:id", controller.DeleteGallery(db.DB))
	}
}
