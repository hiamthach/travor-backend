package dto

import "github.com/travor-backend/model"

type PackageTable struct{}

func (*PackageTable) TableName() string {
	return "packages"
}

type PackageDto struct {
	model.Package
	Destination model.Destination `json:"destination" gorm:"foreignKey:DesID"`
}

func (*PackageDto) TableName() string {
	return "packages"
}

type PackageRequestParam struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type PackagesAllResponse struct {
	Data []PackageDto `json:"data"`
	Pagination
}

type PackageRequestBody struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name" binding:"required"`
	Details      string `json:"details" binding:"required"`
	Price        int    `json:"price" binding:"required"`
	DesID        uint64 `json:"des_id" binding:"required"`
	ImgURL       string `json:"img_url"`
	Duration     string `json:"duration"`
	NumberPeople int    `json:"number_people"`
	Types        []struct {
		ID uint64 `json:"id"`
	} `json:"types"`
	PackageTable
}

type PackageRequestUpdateBody struct {
	Name         string `json:"name"`
	Details      string `json:"details"`
	Price        int    `json:"price"`
	DesID        uint64 `json:"des_id"`
	ImgURL       string `json:"img_url"`
	Duration     string `json:"duration"`
	NumberPeople int    `json:"number_people"`
	Types        []struct {
		ID uint64 `json:"id"`
	} `json:"types"`
	PackageTable
}
