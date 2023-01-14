package main

import (
	"log"

	"github.com/Watson-Sei/ddd-api-golang/application/handlers"
	usecases "github.com/Watson-Sei/ddd-api-golang/application/usecases"
	"github.com/Watson-Sei/ddd-api-golang/domain/repositories"
	"github.com/Watson-Sei/ddd-api-golang/infrastructure/database/gorm"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := gorm.NewGormDB(":memory:")
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	e := echo.New()

	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(userUsecase)

	e.GET("/users", userHandler.GetUsers)
	e.POST("/users", userHandler.CreateUser)

	e.Logger.Fatal(e.Start(":1323"))
}
