package models

import (
	"time"
)

type Post struct {
	Id          uint   `gorm:";primaryKey;AutoIncrement"`
	TextContent string `gorm:"type:varchar(255)"`
	UserId      uint   `gorm:"foreignKey:UserRefer"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User `gorm:"foreignKey:UserId"`
}
