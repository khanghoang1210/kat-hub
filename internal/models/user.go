package models

import (
	"time"
)

type User struct {
	Id        uint `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	UserName  string
	FullName  string
	Email     string
	Password  string
	AvatarUrl string
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
