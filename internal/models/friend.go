package models

import "time"

type Friend struct {
	Id        uint `gorm:"primaryKey;autoIncrement"`
	UserId    uint `gorm:"not null"`
	FriendId  uint `gorm:"not null"`
	User      User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Friend    User `gorm:"foreignKey:FriendId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
}
