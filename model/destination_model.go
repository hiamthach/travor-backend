package model

type Destination struct {
	ID          uint64 `json:"id" gorm:"id"`
	Name        string `json:"name" gorm:"name"`
	Description string `json:"description" gorm:"description"`
	Price       int    `json:"price" gorm:"price"`
	Country     string `json:"country" gorm:"country"`
	VisaRequire string `json:"visa_require" gorm:"visa_require"`
	Language    string `json:"language" gorm:"language"`
	Currency    string `json:"currency" gorm:"currency"`
	Area        string `json:"area" gorm:"area"`
	Location    string `json:"location" gorm:"location"`
}

func (*Destination) TableName() string {
	return "destinations"
}
