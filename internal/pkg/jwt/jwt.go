package token

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/akekawich36/go-authen/configs"
	"github.com/akekawich36/go-authen/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenService interface {
	GenerateAccessToken(user *models.User) (string, error)
	GenerateRefreshToken(user *models.User) (string, error)
	ValidateAccessToken(tokenString string) (*TokenClaims, error)
	ValidateRefreshToken(tokenString string) (*TokenClaims, error)
}

type jwtService struct {
	accessSecret       string
	refreshSecret      string
	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
}

func NewJWTService() *jwtService {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	return &jwtService{
		accessSecret:       config.JWT.Secret,
		refreshSecret:      config.JWT.Secret,
		accessTokenExpiry:  config.JWT.AccessTokenExpiry,
		refreshTokenExpiry: config.JWT.RefreshTokenExpiry,
	}
}

func (s *jwtService) GenerateAccessToken(user *models.User) (string, error) {
	claims := TokenClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.accessTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.accessSecret))
}

func (s *jwtService) GenerateRefreshToken(user *models.User) (string, error) {
	claims := TokenClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.accessTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.accessSecret))
}

func (s *jwtService) ValidateAccessToken(tokenString string) (*TokenClaims, error) {
	token, error := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.accessSecret), nil
	})

	if error != nil {
		return nil, error
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (s *jwtService) ValidateRefreshToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.refreshSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
