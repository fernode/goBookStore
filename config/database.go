package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, nil
}
