package dto

// DestinationDTO is a model that client use when updating a book
type DestinationDTO struct {
	ID          uint64 `json:"id" form:"id"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Price       int    `json:"price" form:"price" binding:"required"`
	Country     string `json:"country" form:"country" binding:"required"`
	VisaRequire string `json:"visa_require" form:"visa_require"`
	Language    string `json:"language" form:"language"`
	Currency    string `json:"currency" form:"currency"`
	Area        string `json:"area" form:"area"`
	Location    string `json:"location" form:"location"`
}
