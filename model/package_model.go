package model

// Package struct
type Package struct {
	ID           uint64 `json:"id" gorm:"primaryKey"`
	DesID        uint64 `json:"des_id" gorm:"des_id"`
	Name         string `json:"name" gorm:"name"`
	Details      string `json:"details" gorm:"details"`
	Price        int32  `json:"price" gorm:"price"`
	ImgURL       string `json:"img_url" gorm:"img_url"`
	Duration     string `json:"duration" gorm:"duration"`
	NumberPeople int32  `json:"number_people" gorm:"number_people"`
	Types        []Type `gorm:"many2many:packages_types;joinForeignKey:PID;JoinReferences:TID" json:"types"`
}

// TableName for package
func (*Package) TableName() string {
	return "packages"
}
