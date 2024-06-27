package responses

import "time"

type UserResponse struct {
	Id        uint
	UserName  string
	FullName  string
	Email     string
	Gender    string
	AvatarUrl string
	CreatedAt time.Time
}
