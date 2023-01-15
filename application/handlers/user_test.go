package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Watson-Sei/ddd-api-golang/application/usecases"
	"github.com/Watson-Sei/ddd-api-golang/domain/models"
	"github.com/Watson-Sei/ddd-api-golang/domain/repositories"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	db, _ := gorm.Open("sqlite3", ":memory:")
	defer db.Close()
	db.AutoMigrate(&models.User{})
	db.Create(&models.User{
		ID:        1,
		FirstName: "Watson",
		LastName:  "Sei",
		Email:     "seinabehack@gmail.com",
	})

	userUsecase := usecases.NewUserUsecase(repositories.NewUserRepository(db))
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewUserHandler(userUsecase)
	if assert.NoError(t, h.GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var users []models.User
		err := json.Unmarshal(rec.Body.Bytes(), &users)
		require.NoError(t, err)
		assert.Len(t, users, 1)
	}
}

func TestCreateUser(t *testing.T) {
	db, _ := gorm.Open("sqlite3", ":memory:")
	defer db.Close()
	db.AutoMigrate(&models.User{})

	user := models.User{
		FirstName: "Watson",
		LastName:  "Sei",
		Email:     "seinabehack@gmail.com",
	}
	db.Create(&user)

	userUsecase := usecases.NewUserUsecase(repositories.NewUserRepository(db))
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte(`{"id":3,"first_name":"Mary","last_name":"Jane","email":"maryjane@gmail.com"}`)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewUserHandler(userUsecase)
	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var user models.User
		err := json.Unmarshal(rec.Body.Bytes(), &user)
		require.NoError(t, err)
		assert.Equal(t, "Mary", user.FirstName)
		assert.Equal(t, "Jane", user.LastName)
		assert.Equal(t, "maryjane@gmail.com", user.Email)
		var count int
		db.Model(&models.User{}).Count(&count)
		assert.Equal(t, 2, count)
	}
}
