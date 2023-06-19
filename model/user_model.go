package model

type User struct {
	Username          string `json:"username" gorm:"primaryKey"`
	HashedPassword    string `json:"hashed_password" gorm:"hashed_password"`
	FullName          string `json:"full_name" gorm:"full_name"`
	Email             string `json:"email" gorm:"email"`
	Phone             string `json:"phone" gorm:"phone"`
	PasswordChangedAt string `json:"password_changed_at" gorm:"password_changed_at"`
	CreatedAt         string `json:"created_at" gorm:"created_at"`
	Role              string `json:"role" gorm:"role"`
}

func (*User) TableName() string {
	return "users"
}
