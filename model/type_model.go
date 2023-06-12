package model

type Type struct {
	Id       int64     `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name" gorm:"name"`
	Packages []Package `gorm:"many2many:packages_types;joinForeignKey:TID;JoinReferences:PID" json:"packages"`
}

func (*Type) TableName() string {
	return "types"
}
