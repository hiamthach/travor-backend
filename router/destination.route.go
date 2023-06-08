package router

import "github.com/gin-gonic/gin"

func destinationRoutes(router *gin.RouterGroup) {
	destinations := router.Group("/destinations")
	{
		destinations.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Test",
			})
		})
	}
}
