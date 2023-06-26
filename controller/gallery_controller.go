package controller

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/travor-backend/docs"
	"github.com/travor-backend/model"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

// @Summary Get galleries by destination ID
// @Description Retrieves a list of galleries based on the provided destination ID
// @Tags Galleries
// @Accept json
// @Produce json
// @Param des path int true "Destination ID"
// @Success 200 {object} model.Gallery
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /galleries/{des} [get]
func GetGalleriesByDesID(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		desID := ctx.Param("des")

		var galleries []model.Gallery
		if err := db.Where("des_id = ?", desID).Find(&galleries).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, util.SuccessResponse("Get galleries successfully", galleries))
	}
}

// @Summary Upload image to gallery
// @Description Uploads an image to the gallery of a specific destination
// @Tags Galleries
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Param des path int true "Destination ID"
// @Param image formData file true "Image file to upload"
// @Success 200 {object} model.Gallery
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /galleries/{des}/upload [post]
func UploadToGallery(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, err := ctx.FormFile("image")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}
		desIDStr := ctx.Param("des")
		desID, err := strconv.ParseUint(desIDStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}
		var gallery model.Gallery

		url, err := util.UploadFile(file, ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		gallery.DesID = desID
		gallery.URL = url.MediaLink

		if err = db.Create(&gallery).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, util.SuccessResponse("Upload image successfully", gallery))
	}
}

func UpdateGallery(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// @Summary Delete a gallery
// @Description Deletes a gallery by ID
// @Tags Galleries
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {access_token}" default(Bearer <access_token>)
// @Param id path int true "Gallery ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /galleries/{id} [delete]
func DeleteGallery(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var gallery model.Gallery
		if err := db.Where("id = ?", id).Delete(&gallery).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, util.SuccessResponse("Delete gallery successfully", nil))
	}
}
