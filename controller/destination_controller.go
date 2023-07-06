package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/travor-backend/docs"
	"github.com/travor-backend/dto"
	"github.com/travor-backend/model"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

// @Summary Get destinations
// @Description Retrieves a list of destinations
// @Tags Destinations
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} dto.DestinationsAllResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /destinations [get]
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

// @Summary Get a destination by ID
// @Description Retrieves a destination based on the provided ID
// @Tags Destinations
// @Accept json
// @Produce json
// @Param id path int true "Destination ID"
// @Success 200 {object} model.Destination
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /destinations/{id} [get]
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

// @Summary Create a destination
// @Description Creates a new destination
// @Tags Destinations
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Param body body dto.DestinationRequestBody true "Destination object to create"
// @Success 201 {object} dto.DestinationRequestBody
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /destinations [post]
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

// @Summary Update a destination
// @Description Updates an existing destination
// @Tags Destinations
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Param id path int true "Destination ID"
// @Param destination body dto.DestinationRequestUpdateBody true "Updated destination object"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /destinations/{id} [put]
func UpdateDestination(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var destination dto.DestinationRequestUpdateBody
		idStr := ctx.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		err = ctx.ShouldBindJSON(&destination)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		result := db.Model(&dto.DestinationRequestUpdateBody{}).Where("id = ?", id).Updates(&destination)
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
		ctx.JSON(http.StatusOK, util.SuccessResponse("Destination updated successfully", nil))
	}
}

// @Summary Delete a destination
// @Description Deletes a destination by ID
// @Tags Destinations
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Param id path int true "Destination ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /destinations/{id} [delete]
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
