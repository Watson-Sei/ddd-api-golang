package gorm

import (
	"github.com/Watson-Sei/ddd-api-golang/domain/models"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	Create(user *models.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) FindAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}
