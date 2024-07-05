package responses

import "time"

type PostResponse struct {
	Id          uint         `json:"id"`
	TextContent string       `json:"textContent"`
	Author      UserResponse `json:"author"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}