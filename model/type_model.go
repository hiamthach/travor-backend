package model

type Type struct {
	Id   int64  `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"`
}

func (*Type) TableName() string {
	return "types"
}
