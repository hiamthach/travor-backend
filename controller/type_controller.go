package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/travor-backend/dto"
	"github.com/travor-backend/model"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

func GetTypes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var types []model.Type
		db.Find(&types)
		c.JSON(http.StatusOK, types)
	}
}

func GetTypeById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var type_ model.Type
		id := c.Param("id")
		result := db.First(&type_, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, util.ErrorResponse(result.Error))
			return
		}
		c.JSON(http.StatusOK, type_)
	}
}

func CreateType(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var type_ dto.TypeDto
		err := c.ShouldBindJSON(&type_)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		result := db.Create(&type_)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, util.ErrorResponse(result.Error))
			return
		}
		c.JSON(http.StatusOK, type_)
	}
}

func UpdateType(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var type_ dto.TypeDto
		id := c.Param("id")
		db.First(&type_, id)
		err := c.ShouldBindJSON(&type_)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}
		db.Save(&type_)
		c.JSON(http.StatusOK, util.SuccessResponse("Update type successfully", type_))
	}
}

func DeleteType(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var type_ model.Type
		id := c.Param("id")
		result := db.Delete(&type_, id)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, util.ErrorResponse(result.Error))
				return
			} else {
				c.JSON(http.StatusInternalServerError, util.ErrorResponse(result.Error))
				return
			}
		}
		c.JSON(http.StatusOK, util.SuccessResponse("Delete type successfully", nil))
	}
}
