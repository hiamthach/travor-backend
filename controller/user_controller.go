package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/travor-backend/docs"
	"github.com/travor-backend/dto"

	// "github.com/travor-backend/model"

	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

// @Summary Get all users
// @Description Retrieves a list of all users
// @Tags Users
// @Produce json
// @Success 200 {array} dto.UserDto
// @Failure 500 {object} model.ErrorResponse
// @Router /users [get]
func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var users []dto.UserDto

		if err := db.Find(&users).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, users)
	}
}

// @Summary Update user info
// @Description Updates the information of a user
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Param username path string true "Username"
// @Param user body dto.UpdateUserInfo true "Updated user information"
// @Success 200 {object} model.SuccessResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /users/{username} [put]
func UpdateUserInfo(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dto.UpdateUserInfo
		username := ctx.Param("username")
		var user dto.UserDto

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		if err := db.Model(&user).Where("username = ?", username).Updates(body).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, util.SuccessResponse("Update user info successfully", nil))
	}
}
