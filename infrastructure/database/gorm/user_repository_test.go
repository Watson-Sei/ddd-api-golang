package gorm

import (
	"testing"

	"github.com/Watson-Sei/ddd-api-golang/domain/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestUserRepo_FindAll(t *testing.T) {
	// mock db
	db, _ := gorm.Open("sqlite3", ":memory:")
	defer db.Close()
	db.AutoMigrate(&models.User{})

	userRepo := NewUserRepository(db)

	// insert test data
	user := models.User{
		ID:        1,
		FirstName: "Watson",
		LastName:  "Sei",
		Email:     "seinabehack@gmail.com",
	}
	db.Create(&user)

	// run test
	users, err := userRepo.FindAll()

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
	assert.Equal(t, "Watson", users[0].FirstName)
	assert.Equal(t, "Sei", users[0].LastName)
	assert.Equal(t, "seinabehack@gmail.com", users[0].Email)
}

func TestUserRepo_Create(t *testing.T) {
	// mock db
	db, _ := gorm.Open("sqlite3", ":memory:")
	defer db.Close()
	db.AutoMigrate(&models.User{})

	userRepo := NewUserRepository(db)

	// run test
	user := models.User{
		ID:        1,
		FirstName: "Watson",
		LastName:  "Sei",
		Email:     "seinabehack@gmail.com",
	}
	err := userRepo.Create(&user)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, "Watson", user.FirstName)
	assert.Equal(t, "Sei", user.LastName)
	assert.Equal(t, "seinabehack@gmail.com", user.Email)
}
