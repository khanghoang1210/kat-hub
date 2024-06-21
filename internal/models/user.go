package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uuid.UUID    			`gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserName  string				
	FullName  string
	Email     string
	Password  string
	AvatarUrl string
	Gender 	  string
	CreatedAt time.Time
	UpdatedAt time.Time
}