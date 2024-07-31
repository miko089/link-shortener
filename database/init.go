package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("links.db"), &gorm.Config{})
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Link{})
}