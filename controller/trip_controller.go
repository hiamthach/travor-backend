package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/travor-backend/dto"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

func GetTrips(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response []dto.TripDto

		if err := db.Find(&response).Error; err != nil {
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}

func GetTripById(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response dto.TripDto
		id := ctx.Param("id")
		if err := db.First(&response, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}

func CreateTrip(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dto.TripDtoBody
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		trip := dto.TripDto{
			User:      body.User,
			PId:       body.PId,
			Total:     body.Total,
			Notes:     body.Notes,
			StartDate: body.StartDate,
		}

		if err := db.Create(&trip).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, trip)
	}
}

func UpdateTrip(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dto.TripDtoUpdate
		idStr := ctx.Param("id")
		id, err := util.StringToUInt64(idStr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		trip := dto.TripDto{
			Id:        id,
			Total:     body.Total,
			Notes:     body.Notes,
			StartDate: body.StartDate,
		}

		result := db.Model(&dto.TripDto{}).Where("id = ?", id).Updates(&trip)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		// Check if no rows were affected
		if result.RowsAffected == 0 {
			err := gorm.ErrRecordNotFound
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, util.SuccessResponse("Successfully updated trip", trip))
	}
}

func DeleteTrip(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := util.StringToUInt64(idStr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		trip := dto.TripDto{
			Id: id,
		}

		result := db.Delete(&trip)

		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		// Check if no rows were affected
		if result.RowsAffected == 0 {
			err := gorm.ErrRecordNotFound
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, util.SuccessResponse("Successfully deleted trip", trip))
	}
}
