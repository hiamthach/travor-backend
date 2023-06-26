package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	"github.com/travor-backend/constant"
	_ "github.com/travor-backend/docs"
	"github.com/travor-backend/dto"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

// @Summary Get trips
// @Description Retrieves a list of trips
// @Tags Trips
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Success 200 {array} dto.TripDto
// @Failure 404 {object} model.ErrorResponse
// @Router /trips [get]
func GetTrips(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response []dto.TripDto
		payload := ctx.MustGet(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)

		if result := db.Where("trips.user = ?", payload.Username).Find(&response); result.Error != nil {
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(result.Error))
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}

// @Summary Get trip by ID
// @Description Retrieves a trip by its ID
// @Tags Trips
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Param id path int true "Trip ID"
// @Success 200 {object} dto.TripDto
// @Failure 404 {object} model.ErrorResponse
// @Router /trips/{id} [get]
func GetTripById(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response dto.TripDto
		id := ctx.Param("id")
		payload := ctx.MustGet(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
		if err := db.Where("trips.user = ?", payload.Username).First(&response, id).Error; err != nil {
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

// @Summary Create a new trip
// @Description Creates a new trip
// @Tags Trips
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Param body body dto.TripDtoBody true "Trip object to be created"
// @Success 200 {object} dto.TripDto
// @Failure 500 {object} model.ErrorResponse
// @Router /trips [post]
func CreateTrip(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dto.TripDtoBody
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		payload := ctx.MustGet(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)

		trip := dto.TripDto{
			User:      payload.Username,
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

// @Summary Update a trip
// @Description Updates an existing trip
// @Tags Trips
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Param id path string true "Trip ID"
// @Param body body dto.TripDtoUpdate true "Updated trip object"
// @Success 200 {object} model.SuccessResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /trips/{id} [put]
func UpdateTrip(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dto.TripDtoUpdate
		idStr := ctx.Param("id")
		id, err := util.StringToUInt64(idStr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}
		payload := ctx.MustGet(constant.AUTHORIZATION_PAYLOAD_KEY).(*util.Payload)
		role := ctx.MustGet(constant.AUTHORIZATION_ROLE).(int)

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

		var result *gorm.DB
		if role == constant.USER_ROLE {
			result = db.Model(&dto.TripDto{}).Where("id = ? AND trips.user = ?", id, payload.Username).Updates(&trip)
		} else {
			result = db.Model(&dto.TripDto{}).Where("id = ?", id).Updates(&trip)
		}

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

// @Summary Delete a trip
// @Description Deletes a trip by ID
// @Tags Trips
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Param id path string true "Trip ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /trips/{id} [delete]
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
