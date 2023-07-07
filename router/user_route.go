package router

import (
	"github.com/gin-gonic/gin"
	"github.com/travor-backend/controller"
	"github.com/travor-backend/db"
	"github.com/travor-backend/middleware"
)

func userRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.GET("/", controller.GetUsers(db.DB))

		users.Use(middleware.AuthMiddleware())
		users.Use(middleware.AdminMiddleware(db.DB))
		users.PUT("/info/:username", controller.UpdateUserInfo(db.DB))
	}
}
