package responses

import "time"

type CommentResponse struct {
	Id        uint         `json:"id"`
	Content   string       `json:"content"`
	Author    UserResponse `json:"author"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}
