package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/travor-backend/dto"
	"github.com/travor-backend/util"
	"gorm.io/gorm"
)

// GetPackages
func GetPackages(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response dto.PackagesAllResponse
		var params dto.PackageRequestParam

		if err := ctx.ShouldBindQuery(&params); err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		// Load the destination data from the database
		err := db.Preload("Destination").Limit(params.PageSize).Offset(params.PageSize * (params.Page - 1)).Find(&response.Data).Error

		db.Table("packages").Count(&response.Total)
		response.CurrentPage = params.Page
		response.PageSize = params.PageSize

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}

// GetPackageById
func GetPackageById(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var packageModel dto.PackageDto
		id := ctx.Param("id")
		err := db.Preload("Destination").First(&packageModel, id).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, packageModel)
	}
}

// CreatePackage
func CreatePackage(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var packageModel dto.PackageRequestBody
		err := ctx.ShouldBindJSON(&packageModel)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		err = db.Create(&packageModel).Error

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusCreated, packageModel)
	}
}

// UpdatePackage
func UpdatePackage(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var packageModel dto.PackageRequestBody
		idStr := ctx.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		err = ctx.ShouldBindJSON(&packageModel)

		packageModel.ID = id
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		result := db.Model(&dto.PackageRequestBody{}).Where("id = ?", id).Updates(&packageModel)

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

		ctx.JSON(http.StatusOK, packageModel)
	}
}

// DeletePackage
func DeletePackage(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var packageModel dto.PackageDto
		idStr := ctx.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		result := db.Delete(&packageModel, id)

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

		ctx.JSON(http.StatusOK, util.SuccessResponse("Package deleted successfully", nil))
	}
}
