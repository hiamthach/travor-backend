package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/travor-backend/dto"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

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

func GetUserByUsername(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user dto.UserDto
		username := ctx.Param("username")

		if result := db.Where("username = ?", username).First(&user); result.Error != nil {
			err := result.Error
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
				return
			}

			ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dto.UserRegisterReq
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		hashedPassword, err := util.HashPassword(body.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		arg := dto.CreateUserParams{
			Username:       body.Username,
			Email:          body.Email,
			Phone:          body.Phone,
			FullName:       body.FullName,
			HashedPassword: hashedPassword,
		}

		if err := db.Create(&arg).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, body)
	}
}

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
