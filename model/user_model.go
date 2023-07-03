package model

import "time"

type User struct {
	Username          string    `json:"username" gorm:"primaryKey"`
	HashedPassword    string    `json:"hashed_password" gorm:"hashed_password"`
	FullName          string    `json:"full_name" gorm:"full_name"`
	Email             string    `json:"email" gorm:"email"`
	Phone             string    `json:"phone" gorm:"phone"`
	PasswordChangedAt time.Time `json:"password_changed_at" gorm:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at" gorm:"created_at"`
	Role              int       `json:"role" gorm:"role"`
	Status            bool      `json:"status" gorm:"status"`
}

func (*User) TableName() string {
	return "users"
}
