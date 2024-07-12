package repository

import (
	"kathub/internal/models"
	"kathub/pkg/requests"
)

type PostRepository interface {
	Create(post *requests.CreatePostReq, currentUser uint) (bool, error)
	GetAll() ([]*models.Post, error)
	Update(post *requests.CreatePostReq) (bool, error)
}