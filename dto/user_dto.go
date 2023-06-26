package dto

import (
	"time"

	uuid "github.com/google/uuid"
)

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
	Status            bool   `json:"status"`
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
	Role           int    `json:"role" binding:"required"`
	Status         bool   `json:"status" binding:"required"`
	UserTable
}

type UpdateUserInfo struct {
	Phone    string `json:"phone"`
	FullName string `json:"full_name"`
	Status   bool   `json:"status"`
	UserTable
}

type UserLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserTable
}

type UserLoginRes struct {
	SessionID             uuid.UUID `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	User                  UserDto   `json:"user"`
	UserTable
}

type UserSessionParams struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	ExpiresAt    time.Time `json:"expires_at"`
	IsBlocked    bool      `json:"is_blocked"`
}

type RenewTokenReq struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RenewTokenRes struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (*UserSessionParams) TableName() string {
	return "sessions"
}
