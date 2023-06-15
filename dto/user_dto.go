package dto

type UserTable struct{}

func (*UserTable) TableName() string {
	return "users"
}

type UserDto struct {
	Username          string `json:"username"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	FullName          string `json:"full_name"`
	CreatedAt         string `json:"created_at"`
	PasswordChangedAt string `json:"password_changed_at"`
	UserTable
}

func (u *UserDto) TableName() string {
	return "users"
}

type UserRegisterReq struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserTable
}

type CreateUserParams struct {
	Username       string `json:"username" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Phone          string `json:"phone" binding:"required"`
	FullName       string `json:"full_name" binding:"required"`
	HashedPassword string `json:"hashed_password" binding:"required"`
	UserTable
}
