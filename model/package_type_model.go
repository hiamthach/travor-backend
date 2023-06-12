package model

type PackageType struct {
	PID uint64 `json:"p_id" gorm:"p_id"`
	TID uint64 `json:"t_id" gorm:"t_id"`
}

func (*PackageType) TableName() string {
	return "packages_types"
}
