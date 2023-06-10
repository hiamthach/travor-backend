package model

// Package struct
type Package struct {
	ID           uint64 `json:"id" gorm:"id"`
	DesID        uint64 `json:"des_id" gorm:"des_id"`
	Name         string `json:"name" gorm:"name"`
	Details      string `json:"details" gorm:"details"`
	Price        int    `json:"price" gorm:"price"`
	ImgURL       string `json:"img_url" gorm:"img_url"`
	Duration     string `json:"duration" gorm:"duration"`
	NumberPeople int    `json:"number_people" gorm:"number_people"`
}

// TableName for package
func (*Package) TableName() string {
	return "packages"
}