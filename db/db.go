package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"topinambur02.com/m/v2/model"
)

func InitDB() gorm.DB {
	sqliteConnection := sqlite.Open("database.db")
	db, err := gorm.Open(sqliteConnection, &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&model.User{})

	if err != nil {
		panic("Migration failed: " + err.Error())
	}

	return *db
}
