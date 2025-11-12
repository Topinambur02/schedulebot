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
		Username:  config.ADMIN_USERNAME,
		Role:      "ADMIN",
		CreatedAt: time.Now(),
		Tg_id:     config.TG_ID,
	}
	db.Where(model.User{Tg_id: config.TG_ID}).FirstOrCreate(&user)

	return &Database{Db: db}
}

func (d *Database) FindByTgID(tgId int64) (*model.User, error) {
	var user model.User
	result := d.Db.Where("tg_id = ?", tgId).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (d *Database) Create(user *model.User) error {
	return d.Db.Create(&user).Error
}

func (d *Database) Update(user *model.User) error {
	return d.Db.Save(user).Error
}
