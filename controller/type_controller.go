package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/travor-backend/docs"
	"github.com/travor-backend/dto"
	"github.com/travor-backend/model"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

// @Summary Get types
// @Description Retrieves a list of types
// @Tags Types
// @Accept json
// @Produce json
// @Success 200 {array} model.Type
// @Failure 500 {object} model.ErrorResponse
// @Router /types [get]
func GetTypes(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var types []model.Type
		db.Find(&types)
		ctx.JSON(http.StatusOK, types)
	}
}

// @Summary Get type by ID
// @Description Retrieves a type by its ID
// @Tags Types
// @Accept json
// @Produce json
// @Param id path int true "Type ID"
// @Success 200 {object} model.Type
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /types/{id} [get]
func GetTypeById(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var type_ model.Type
		id := ctx.Param("id")
		result := db.First(&type_, id)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(result.Error))
			return
		}
		ctx.JSON(http.StatusOK, type_)
	}
}

// @Summary Create new type
// @Description Creates a new type
// @Tags Types
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Param type_ body dto.TypeDto true "Type object to create"
// @Success 200 {object} dto.TypeDto
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /types [post]
func CreateType(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var type_ dto.TypeDto
		err := ctx.ShouldBindJSON(&type_)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		result := db.Create(&type_)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(result.Error))
			return
		}
		ctx.JSON(http.StatusOK, type_)
	}
}

// @Summary Update type by ID
// @Description Updates an existing type by ID
// @Tags Types
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Param id path int true "Type ID"
// @Param type_ body dto.TypeDto true "Type object to update"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /types/{id} [put]
func UpdateType(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var type_ dto.TypeDto
		id := ctx.Param("id")
		db.First(&type_, id)
		err := ctx.ShouldBindJSON(&type_)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}
		db.Save(&type_)
		ctx.JSON(http.StatusOK, util.SuccessResponse("Update type successfully", type_))
	}
}

// @Summary Delete type by ID
// @Description Deletes a type by ID
// @Tags Types
// @Param id path int true "Type ID"
// @Produce json
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Success 200 {object} model.SuccessResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /types/{id} [delete]
func DeleteType(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var type_ model.Type
		id := ctx.Param("id")
		result := db.Delete(&type_, id)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, util.ErrorResponse(result.Error))
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(result.Error))
				return
			}
		}
		ctx.JSON(http.StatusOK, util.SuccessResponse("Delete type successfully", nil))
	}
}
