package repositories

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

func (u *userRepo) FindAll() ([]models.User, error) {
	var users []models.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepo) Create(user *models.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
