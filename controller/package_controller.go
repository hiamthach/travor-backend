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

// @Summary Get packages
// @Description Retrieves a list of packages
// @Tags Packages
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {array} dto.PackagesAllResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /packages [get]
func GetPackages(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response dto.PackagesAllResponse
		var params dto.PackageRequestParam

		if err := ctx.ShouldBindQuery(&params); err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		// Load the destination data from the database
		err := db.Preload("Types").Preload("Destination").Limit(params.PageSize).Offset(params.PageSize * (params.Page - 1)).Find(&response.Data).Error

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

// @Summary Get a package by ID
// @Description Retrieves a package based on the provided ID
// @Tags Packages
// @Accept json
// @Produce json
// @Param id path int true "Package ID"
// @Success 200 {object} dto.PackageDto
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /packages/{id} [get]
func GetPackageById(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var packageModel dto.PackageDto
		id := ctx.Param("id")
		err := db.Preload("Types").Preload("Destination").First(&packageModel, id).Error

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

// @Summary Create a package
// @Description Creates a new package
// @Tags Packages
// @Accept json
// @Produce json
// @Param body body dto.PackageRequestBody true "Package object to create"
// @Success 201 {object} model.Package
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /packages [post]
func CreatePackage(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dto.PackageRequestBody
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		// Map the request body to the model
		packageCreated := model.Package{
			Name:         body.Name,
			Details:      body.Details,
			Price:        body.Price,
			DesID:        body.DesID,
			ImgURL:       body.ImgURL,
			Duration:     body.Duration,
			NumberPeople: body.NumberPeople,
		}

		if err := db.Create(&packageCreated).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		// Loop through the types and create the package type
		for _, typ := range body.Types {
			if err := db.Create(&model.PackageType{
				PID: packageCreated.ID,
				TID: typ.ID,
			}).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusCreated, packageCreated)
	}
}

// @Summary Update a package
// @Description Updates an existing package
// @Tags Packages
// @Accept json
// @Produce json
// @Param id path int true "Package ID"
// @Param body body dto.PackageRequestBody true "Updated package object"
// @Success 200 {object} dto.PackageRequestBody
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /packages/{id} [put]
func UpdatePackage(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body dto.PackageRequestBody
		idStr := ctx.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		body.ID = id
		if err = ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		var packageUpdated model.Package
		if err := db.First(&packageUpdated, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		packageUpdated = model.Package{
			ID:           body.ID,
			Name:         body.Name,
			Details:      body.Details,
			Price:        body.Price,
			DesID:        body.DesID,
			ImgURL:       body.ImgURL,
			Duration:     body.Duration,
			NumberPeople: body.NumberPeople,
		}

		if err := db.Save(&packageUpdated).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		// Delete all existing package types associated with the package
		if err := db.Delete(model.PackageType{}, "p_id = ?", packageUpdated.ID).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
			return
		}

		// Create and save the new package types
		for _, typ := range body.Types {
			if err := db.Create(&model.PackageType{
				PID: packageUpdated.ID,
				TID: typ.ID,
			}).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusOK, body)
	}
}

// @Summary Delete a package
// @Description Deletes a package by ID
// @Tags Packages
// @Accept json
// @Produce json
// @Param id path int true "Package ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 404 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /packages/{id} [delete]
func DeletePackage(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var packageModel dto.PackageDto
		idStr := ctx.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
			return
		}

		if err := db.Delete(model.PackageType{}, "p_id = ?", id).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
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
