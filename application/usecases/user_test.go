package usecases

import (
	"testing"

	"github.com/Watson-Sei/ddd-api-golang/domain/models"
	"github.com/Watson-Sei/ddd-api-golang/domain/repositories"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestUserUsecase(t *testing.T) {
	db, _ := gorm.Open("sqlite3", ":memory:")
	defer db.Close()
	db.AutoMigrate(&models.User{})

	// Test Getusers method
	users := []models.User{{
		ID:        1,
		FirstName: "Watson",
		LastName:  "Sei",
		Email:     "seinabehack@gmail.com",
	}, {
		ID:        2,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@gmail.com",
	}}
	for _, user := range users {
		db.Create(user)
	}

	userRepo := repositories.NewUserRepository(db)
	userUsecase := NewUserUsecase(userRepo)
	users, err := userUsecase.GetUsers()
	if err != nil {
		t.Errorf("Error in GetUsers method: %s", err)
	}
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
	if users[0].FirstName != "Watson" && users[0].LastName != "Sei" {
		t.Errorf("Expected user with FirstName: Watson, LastName: Sei")
	}
	if users[1].FirstName != "John" && users[1].LastName != "Doe" {
		t.Errorf("Expected user with FirstName: John, LastName: Doe")
	}

	// Test CreateUser method
	newUser := &models.User{
		ID:        3,
		FirstName: "Mary",
		LastName:  "Jane",
		Email: "maryjane@gmail.com",
	}
	err = userUsecase.CreateUser(newUser)
	if err != nil {
		t.Errorf("Error in CreateUser method: %s", err)
	}
}
