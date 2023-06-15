package router

import (
	"github.com/gin-gonic/gin"
	"github.com/travor-backend/controller"
	"github.com/travor-backend/db"
)

func userRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.GET("/", controller.GetUsers(db.DB))
		users.GET("/:username", controller.GetUserByUsername(db.DB))
		users.PUT("/", controller.CreateUser(db.DB))
		users.POST("/info/:username", controller.UpdateUserInfo(db.DB))
	}
}