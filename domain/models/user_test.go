package models_test

import (
	"testing"

	"github.com/Watson-Sei/ddd-api-golang/domain/models"
)

func TestUser_TableName(t *testing.T) {
	user := models.User{}
	want := "users"

	if got := user.TableName(); got != want {
		t.Errorf("User.TableName() = %v, want %v", got, want)
	}
}
