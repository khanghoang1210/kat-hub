package models

import "time"

type Like struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	UserId    uint `gorm:"not null"`
	PostId    uint `gorm:"not null"`
	CreatedAt time.Time
}