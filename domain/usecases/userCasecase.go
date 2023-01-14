package usecases

import (
	"github.com/Watson-Sei/ddd-api-golang/domain/models"
	"github.com/Watson-Sei/ddd-api-golang/domain/repositories"
)

type GetUsersUsecase struct {
	userRepo repositories.UserRepository
}

func (u *GetUsersUsecase) Execute() ([]models.User, error) {
	users, err := u.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

type CreaetUserUsecase struct {
	userRepo repositories.UserRepository
}

func (u *CreaetUserUsecase) Execute(user *models.User) error {
	return u.userRepo.Create(user)
}
