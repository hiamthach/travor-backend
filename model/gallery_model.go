package model

type Gallery struct {
	ID    uint64 `json:"id" gorm:"primaryKey"`
	DesID uint64 `json:"des_id" gorm:"des_id"`
	URL   string `json:"url" gorm:"url"`
}

func (*Gallery) TableName() string {
	return "galleries"
}
