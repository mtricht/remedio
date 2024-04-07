package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB = getDatabase()

func getDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Medication{})
	return db
}
