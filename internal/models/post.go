package models

import (
	"time"

	"github.com/lib/pq"
)

type Post struct {
	Id           uint           `gorm:";primaryKey;AutoIncrement"`
	TextContent  string         `gorm:"type:varchar(255)"`
	ImageContent pq.StringArray `gorm:"type:text[]"`
	UserId       uint           `gorm:"foreignKey:UserRefer"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         User      `gorm:"foreignKey:UserId"`
	Likes        []Like    `gorm:"foreignKey:PostId"`
	Comments     []Comment `gorm:"foreignKey:PostId"`
}
