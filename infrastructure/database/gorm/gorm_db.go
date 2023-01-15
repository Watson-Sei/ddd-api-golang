package gorm

import (
	"github.com/Watson-Sei/ddd-api-golang/domain/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewGormDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	return db, nil
}
