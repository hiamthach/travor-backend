package util

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrExpiredToken = errors.New("token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type Maker interface {
	CreateToken(username string, duration time.Duration, privateKey string) (string, *Payload, error)

	VerifyToken(token string, publicKey string) (*Payload, error)
}

type JWTMaker struct {
	secretKey string
}

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssueAt   time.Time `json:"issue_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssueAt:   time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func NewJWTMaker(secretKey string) (Maker, error) {
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration, privateKey string) (string, *Payload, error) {
	now := time.Now().UTC()
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, nil
	}

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", payload, err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", payload, err
	}

	atClaims := make(jwt.MapClaims)
	atClaims["sub"] = username
	atClaims["token_uuid"] = payload.ID
	atClaims["exp"] = payload.ExpiredAt.Unix()
	atClaims["iat"] = now.Unix()
	atClaims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims).SignedString(key)
	if err != nil {
		return "", payload, err
	}

	return token, payload, err
}

func (maker *JWTMaker) VerifyToken(token string, publicKey string) (*Payload, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return nil, err
	}

	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, ErrInvalidToken
		}

		return key, nil
	})
	if err != nil {
		return nil, err
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, ErrInvalidToken
	}

	return &Payload{
		ID:        uuid.MustParse(payload["token_uuid"].(string)),
		Username:  payload["sub"].(string),
		IssueAt:   time.Unix(int64(payload["iat"].(float64)), 0),
		ExpiredAt: time.Unix(int64(payload["exp"].(float64)), 0),
	}, nil
}

func (maker *JWTMaker) RenewToken(refreshToken string, accessTokenDuration, refreshTokenDuration time.Duration) (string, *Payload, error) {
	config, err := LoadConfig(".")
	if err != nil {
		return "", nil, err
	}

	refreshPayload, err := maker.VerifyToken(refreshToken, config.AccessTokenPublicKey)
	if err != nil {
		return "", nil, err
	}

	// Check if the refresh token has expired
	if err := refreshPayload.Valid(); err != nil {
		return "", nil, err
	}

	// Create a new access token using the same username and a new duration
	newAccessToken, _, err := maker.CreateToken(refreshPayload.Username, accessTokenDuration, config.AccessTokenPrivateKey)
	if err != nil {
		return "", nil, err
	}

	return newAccessToken, refreshPayload, nil
}
