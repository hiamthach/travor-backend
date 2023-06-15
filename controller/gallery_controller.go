package controller

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/travor-backend/model"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

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
