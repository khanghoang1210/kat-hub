package models

import (
	"time"
)

type User struct {
	Id        uint   `gorm:";primaryKey;AutoIncrement"`
	UserName  string `gorm:"unique;not null"`
	FullName  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Title     string
	AvatarUrl string
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Post      []Post    `gorm:"foreignKey:UserId"`
	Likes     []Like    `gorm:"foreignKey:UserId"`
	Comments  []Comment `gorm:"foreignKey:UserId"`
}
