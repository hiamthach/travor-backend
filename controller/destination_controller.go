package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/travor-backend/dto"
	"github.com/travor-backend/model"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

// This function retrieves a list of destinations from a database based on pagination params a
func GetDestinations(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response dto.DestinationsAllResponse
		var params dto.DestinationRequestParam
		if err := ctx.ShouldBindQuery(&params); err != nil {
			// Handle the error, if any
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		err := db.Limit(params.PageSize).Offset(params.PageSize * (params.Page - 1)).Find(&response.Data)
		db.Table("destinations").Count(&response.Total)
		response.CurrentPage = params.Page
		response.PageSize = params.PageSize
		if err.Error != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err.Error))
			return
		}
		ctx.JSON(200, response)
	}
}

// This function retrieves a destination by its ID from a database and returns it as a JSON response.
func GetDestinationById(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var destination model.Destination
		id := ctx.Param("id")
		err := db.First(&destination, id)
		if err.Error != nil {
			if err.Error == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, util.ErrorResponse(err.Error))
				return
			}

			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err.Error))
			return
		}
		ctx.JSON(http.StatusOK, destination)
	}
}

// This function creates a destination in a database and returns it as a JSON response.
func CreateDestination(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var destination dto.DestinationRequestBody
		err := ctx.ShouldBindJSON(&destination)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}
		err = db.Create(&destination).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusCreated, destination)
	}
}

// This function updates a destination in a database and returns it as a JSON response.
func UpdateDestination(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var destination dto.DestinationRequestBody
		idStr := ctx.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}
		destination.ID = id
		err = ctx.ShouldBindJSON(&destination)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		result := db.Model(&dto.DestinationRequestBody{}).Where("id = ?", id).Updates(&destination)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(result.Error))
			return
		}

		// Check if no rows were affected
		if result.RowsAffected == 0 {
			err := gorm.ErrRecordNotFound
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, destination)
	}
}

// This function deletes a destination in a database and returns it as a JSON response.
func DeleteDestination(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		result := db.Delete(&model.Destination{}, id)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(result.Error))
			return
		}
		if result.RowsAffected == 0 {
			err := gorm.ErrRecordNotFound
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, util.SuccessResponse("Destination deleted successfully", nil))
	}
}
