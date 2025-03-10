package repository

import (
	"errors"

	"github.com/akekawich36/go-authen/internal/domain/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindById(id uint) (*models.User, error)
	UpdateRefreshToken(userID uint, refreshToken string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository(db)
}

func (r *userRepository) Create(user *models.User) error {
	tx := r.db.Begin()
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ? AND is_deleted = ?", email, false).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) FindByPk(id int) (*models.User, error) {
	var user models.User
	result := r.db.Where("id = ? AND is_deleted = ?", id, false).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) UpdateRefreshToken(userID uint, refreshToken string) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("refresh_token", refreshToken).Error
}
