package models

import "time"

type Like struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	UserId    uint `gorm:"foreignKey:UserRefer"`
	PostId    uint `gorm:"foreignKey:PostRefer"`
	User      User `gorm:"foreignKey:UserId"`
	Post      Post `gorm:"foreignKey:PostId"`
	CreatedAt time.Time
}
