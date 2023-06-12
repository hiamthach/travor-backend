package dto

type TypeDto struct {
	Name string `json:"name" binding:"required"`
}

func (t *TypeDto) TableName() string {
	return "types"
}
