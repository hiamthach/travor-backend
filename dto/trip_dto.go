package dto

import "time"

type TripTable struct{}

func (*TripTable) TableName() string {
	return "trips"
}

type TripDto struct {
	Id        uint64    `json:"id"`
	User      string    `json:"user" binding:"required"`
	PId       int64     `json:"p_id" binding:"required"`
	Total     int64     `json:"total" binding:"required"`
	Notes     string    `json:"notes" `
	StartDate time.Time `json:"start_date" binding:"required"`
	TripTable
}

type TripDtoBody struct {
	User      string    `json:"user" binding:"required"`
	PId       int64     `json:"p_id" binding:"required"`
	Total     int64     `json:"total" binding:"required"`
	Notes     string    `json:"notes"`
	StartDate time.Time `json:"start_date" binding:"required"`
	TripTable
}

type TripDtoUpdate struct {
	Total     int64     `json:"total"`
	Notes     string    `json:"notes"`
	StartDate time.Time `json:"start_date"`
	TripTable
}
