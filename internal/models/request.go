package models

import "time"

type Request struct {
	Id         uint `gorm:"primaryKey;autoIncrement"`
	SenderID   uint `gorm:"foreignKey:UserRefer"`
	ReceiverID uint `gorm:"foreignKey:UserRefer"`
	Status     int  `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Sender     User `gorm:"foreignKey:SenderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Receiver   User `gorm:"foreignKey:ReceiverID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
