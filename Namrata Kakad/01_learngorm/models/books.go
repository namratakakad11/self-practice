package models

import (
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	ID        int     `gorm:"primary-key";autoIncrement json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	publisher *string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
