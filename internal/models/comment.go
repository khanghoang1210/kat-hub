package models

import "time"

type Comment struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	UserId    uint   `gorm:"foreignKey:UserRefer"`
	PostId    uint   `gorm:"foreignKey:PostRefer"`
	Content   string `gorm:"type:text;not null"`
	User      User `gorm:"foreignKey:UserId"`
	Post      Post `gorm:"foreignKey:PostId"`
	CreatedAt time.Time
}