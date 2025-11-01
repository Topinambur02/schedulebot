package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint
	Username  string
	Tg_id     int64
	CreatedAt time.Time
}
