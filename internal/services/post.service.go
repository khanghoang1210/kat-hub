package services

import (
	"io"
	"kathub/pkg/requests"
	"kathub/pkg/responses"
)

type PostService interface {
	Create(req *requests.CreatePostReq, currentUser responses.UserResponse) *responses.ResponseData
	GetAll() *responses.ResponseData
	GetById(id uint) *responses.ResponseData
	Update(req *requests.CreatePostReq, id uint) *responses.ResponseData
	Delete(id uint) *responses.ResponseData
	UploadPostImage(postID int,currentUser responses.UserResponse, data io.Reader, fileNameToUpload string) *responses.ResponseData
	Like(postID int, currentUser responses.UserResponse) *responses.ResponseData
	UnLike(postID int, currentUser responses.UserResponse) *responses.ResponseData
	CreateComment(req *requests.CreateCommentReq, currentUser responses.UserResponse) *responses.ResponseData
	GetCommentsByPostID(postID int) *responses.ResponseData
	UpdateComment(req *requests.UpdateCommentReq, currentUser responses.UserResponse) *responses.ResponseData
}