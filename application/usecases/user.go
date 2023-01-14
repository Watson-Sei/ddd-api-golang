package usecases

import (
	"github.com/Watson-Sei/ddd-api-golang/domain/models"
	"github.com/Watson-Sei/ddd-api-golang/domain/repositories"
)

type UserUsecase interface {
	GetUsers() ([]models.User, error)
	CreateUser(user *models.User) error
}

type userUsecase struct {
	UserRepository repositories.UserRepository
}

func NewUserUsecase(r repositories.UserRepository) UserUsecase {
	return &userUsecase{
		UserRepository: r,
	}
}

func (u *userUsecase) GetUsers() ([]models.User, error) {
	return u.UserRepository.FindAll()
}

func (u *userUsecase) CreateUser(user *models.User) error {
	return u.UserRepository.Create(user)
}
