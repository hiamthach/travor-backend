package model

import "time"

type Trip struct {
	Id        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"user" gorm:"username"`
	PId       uint64    `json:"p_id" gorm:"p_id"`
	Total     int64     `json:"total" gorm:"total"`
	Notes     string    `json:"notes" gorm:"notes"`
	StartDate time.Time `json:"start_date" gorm:"start_date"`
}

func (*Trip) TableName() string {
	return "trips"
}
