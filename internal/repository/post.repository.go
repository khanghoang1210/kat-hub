package repository

import (

	"kathub/pkg/requests"
	"kathub/pkg/responses"
)

type PostRepository interface {
	Create(post *requests.CreatePostReq, currentUser uint) (bool, error)
	GetAll() ([]*responses.PostResponse, error)
	GetById(id uint) (*responses.PostResponse, error)
	Update(req *requests.CreatePostReq, id uint) (bool, error)
	Delete(id uint) (bool, error)
	Like(id uint, user responses.UserResponse)(bool, error)
	UnLike(id uint, user responses.UserResponse)(bool, error)
}