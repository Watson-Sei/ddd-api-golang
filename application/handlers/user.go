package handlers

import (
	"net/http"

	"github.com/Watson-Sei/ddd-api-golang/application/usecases"
	"github.com/Watson-Sei/ddd-api-golang/domain/models"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUsecase usecases.UserUsecase
}

func NewUserHandler(u usecases.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: u,
	}
}

func (u *UserHandler) GetUsers(c echo.Context) error {
	users, err := u.UserUsecase.GetUsers()
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, users)
}

func (u *UserHandler) CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, err)
	}
	if err := u.UserUsecase.CreateUser(&user); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(http.StatusOK, user)
}
