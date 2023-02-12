package models

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// create database folder if not exists
	err := os.MkdirAll("data/database", os.ModePerm)

	db, err := gorm.Open(sqlite.Open("data/database/poll.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Question{}, &Answer{})
	DB = db
}
