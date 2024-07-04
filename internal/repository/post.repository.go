package repository

import (
	"kathub/pkg/requests"
	"kathub/pkg/responses"
)

type PostRepository interface {
	Create(post *requests.CreatePostReq, currentUser uint) (bool, error)
	GetAll() ([]*responses.PostResponse, error)
	Update(post *requests.CreatePostReq) (bool, error)
}