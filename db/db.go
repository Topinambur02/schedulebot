package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"topinambur02.com/m/v2/configs"
	"topinambur02.com/m/v2/model"
)

type Database struct {
	Db *gorm.DB
}

func InitDB() *Database {
	sqliteConnection := sqlite.Open("database.db")
	db, err := gorm.Open(sqliteConnection, &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&model.User{})

	if err != nil {
		panic("Migration failed: " + err.Error())
	}

	config := configs.NewConfig()
	user := model.User{
		Username: config.ADMIN_USERNAME, 
		Role: "ADMIN", 
		CreatedAt: time.Now(), 
		Tg_id: config.TG_ID,
	}
	db.FirstOrCreate(&user)

	return &Database{Db: db}
}

func (d *Database) FindByTgID(tgId int64) *model.User {
	var user model.User
	d.Db.Where("tg_id = ?", tgId).First(&user)
	return &user
}
