package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
)

type Post struct {
	Id          uuid.UUID    		`gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	TextContent string 				`gorm:"type:varchar(255)"`
	User User  						`gorm:"foreignKey:UserRefer"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}