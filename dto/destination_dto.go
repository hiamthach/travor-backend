package dto

import "github.com/travor-backend/model"

type DestinationTable struct{}

func (*DestinationTable) TableName() string {
	return "destinations"
}

// I want to implement the Pagination into the DestinationRequestParam struct
type DestinationRequestParam struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type DestinationsAllResponse struct {
	Data []model.Destination `json:"data"`
	Pagination
}

type DestinationRequestBody struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Country     string `json:"country" binding:"required"`
	VisaRequire string `json:"visa_require"`
	Language    string `json:"language"`
	Currency    string `json:"currency"`
	Area        string `json:"area"`
	Location    string `json:"location"`
	DestinationTable
}
