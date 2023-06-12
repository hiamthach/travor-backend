package model

type Plan struct {
	Id          int64  `json:"id" gorm:"id"`
	PID         int64  `json:"pid" gorm:"pid"`
	Order       uint   `json:"order" gorm:"order"`
	Name        string `json:"name" gorm:"name"`
	Description string `json:"description" gorm:"description"`
}

func (*Plan) TableName() string {
	return "plans"
}
