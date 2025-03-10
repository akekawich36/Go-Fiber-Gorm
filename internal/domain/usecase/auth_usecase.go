package usecase

import (
	"errors"

	"github.com/akekawich36/go-authen/internal/domain/models"
	"github.com/akekawich36/go-authen/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	Register(input models.UserRegister) (*models.AuthResponse, error)
}

type authUseCase struct {
	userRepo     repository.UserRepository
	tokenService token.TokenService
}

func (u *authUseCase) Register(input models.UserRegister) (*models.AuthResponse, error) {
	existingUser, err := u.userRepo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("Email already exists")
	}

	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(input.password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := u.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		User:         *user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
