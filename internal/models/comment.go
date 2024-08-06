package models

import "time"

type Comment struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	UserId    uint   `gorm:"not null"`
	PostId    uint   `gorm:"not null"`
	Content   string `gorm:"type:text;not null"`
	CreatedAt time.Time
}