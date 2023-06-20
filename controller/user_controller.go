package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/travor-backend/docs"
	"github.com/travor-backend/dto"
	"github.com/travor-backend/model"

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

// @Summary Get user by username
// @Description Retrieves a user by their username
// @Tags Users
// @Produce json
// @Param username path string true "Username of the user to retrieve"
// @Success 200 {object} dto.UserDto
// @Failure 404 {object} model.ErrorResponse
// @Router /users/{username} [get]
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

// @Summary Create user
// @Description Creates a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.UserRegisterReq true "User registration details"
// @Success 200 {object} dto.UserRegisterReq
// @Failure 400 {object} model.ErrorResponse
// @Router /users [post]
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
			Role:           1,
			Status:         false,
		}

		if err := db.Create(&arg).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, body)
	}
}

// @Summary Update user info
// @Description Updates the information of a user
// @Tags Users
// @Accept json
// @Produce json
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

// @Summary Login user
// @Description Logs in a user and generates access and refresh tokens
// @Tags Authentication
// @Accept json
// @Produce json
// @Param body body dto.UserLoginReq true "User login request"
// @Success 200 {object} dto.UserLoginRes
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /users/login [post]
func LoginUser(db *gorm.DB, config util.Config, tokenMaker util.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.UserLoginReq
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		var user model.User
		if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		if err := util.CheckPassword(req.Password, user.HashedPassword); err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		accessToken, accessPayload, err := tokenMaker.CreateToken(user.Username, config.AccessTokenDuration)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		refreshToken, refreshTokenPayload, err := tokenMaker.CreateToken(user.Username, config.RefreshTokenDuration)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		session := dto.UserSessionParams{
			ID:           refreshTokenPayload.ID,
			Username:     user.Username,
			RefreshToken: refreshToken,
			UserAgent:    ctx.Request.UserAgent(),
			ClientIp:     ctx.ClientIP(),
			IsBlocked:    false,
			ExpiresAt:    refreshTokenPayload.ExpiredAt,
		}

		if err := db.Create(&session).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		rsp := dto.UserLoginRes{
			SessionID:             session.ID,
			AccessToken:           accessToken,
			RefreshToken:          refreshToken,
			AccessTokenExpiresAt:  accessPayload.ExpiredAt,
			RefreshTokenExpiresAt: refreshTokenPayload.ExpiredAt,
			User: dto.UserDto{
				Username: user.Username,
				Email:    user.Email,
				Phone:    user.Phone,
				FullName: user.FullName,
				Status:   user.Status,
			},
		}

		ctx.JSON(http.StatusOK, rsp)
	}
}
