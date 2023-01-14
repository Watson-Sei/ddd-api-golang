package gorm

import "github.com/jinzhu/gorm"

func NewGormDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
