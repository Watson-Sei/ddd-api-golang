package repositories

import (
	"testing"

	"github.com/Watson-Sei/ddd-api-golang/domain/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestUserRepository(t *testing.T) {
	db, _ := gorm.Open("sqlite3", ":memory:")
	defer db.Close()
	db.AutoMigrate(&models.User{})

	userRepo := NewUserRepository(db)

	// Test FindAll method
	users, err := userRepo.FindAll()
	if err != nil {
		t.Errorf("Error in FindAll method: %s", err)
	}
	if len(users) != 0 {
		t.Errorf("Expected 0 users, got %d", len(users))
	}

	// Test Create method
	newUser := &models.User{
		ID:        1,
		FirstName: "Watson",
		LastName:  "Sei",
		Email:     "seinabehack@gmail.com",
	}
	err = userRepo.Create(newUser)
	if err != nil {
		t.Errorf("Error in Create method: %s", err)
	}
	users, _ = userRepo.FindAll()
	if len(users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(users))
	}
	if users[0].FirstName != "Watson" && users[0].LastName != "Sei" && users[0].Email != "seinabehack@gmail.com" {
		t.Errorf("Expected user with FirstName: Watson, LastName: Sei, Email: seinabehack@gmail.com")
	}
}
