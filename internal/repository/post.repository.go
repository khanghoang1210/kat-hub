package repository

import (
	"kathub/pkg/requests"
	"kathub/pkg/responses"
)

type PostRepository interface {
	Create(post *requests.CreatePostReq, currentUser responses.UserResponse) (*uint, error)
	GetAll() ([]*responses.PostResponse, error)
	GetById(id uint) (*responses.PostResponse, error)
	Update(req *requests.CreatePostReq, id uint) (bool, error)
	Delete(id uint) (bool, error)
	InsertPostImage(postID int, imageUrl string) (bool, error)
	Like(postID int, user responses.UserResponse) (bool, error)
	UnLike(postID int, user responses.UserResponse) (bool, error)
	InsertComment(req *requests.CreateCommentReq, currentUser responses.UserResponse) (bool, error)
	GetAllComment(postID int) ([]*responses.CommentResponse, error)
	UpdateComment(req *requests.UpdateCommentReq, user responses.UserResponse) (bool, error)
	DeleteComment(commentID int, user responses.UserResponse) (bool, error)
}
