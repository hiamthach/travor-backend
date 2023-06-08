package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/travor-backend/model"
	"gorm.io/gorm"
)

func GetDestinations(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var destinations []model.Destination
		err := db.Find(&destinations)
		if err.Error != nil {
			if err.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error,
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error,
			})
			return
		}
		c.JSON(200, gin.H{
			"data": destinations,
		})
	}
}
