package services

import (
	"io"
	"kathub/pkg/requests"
	"kathub/pkg/responses"
)

type PostService interface {
	Create(req *requests.CreatePostReq, currentUser responses.UserResponse,data io.Reader, fileNameToUpload string) *responses.ResponseData
	GetAll() *responses.ResponseData
	GetById(id uint) *responses.ResponseData
	Update(req *requests.CreatePostReq, id uint) *responses.ResponseData
	Delete(id uint) *responses.ResponseData
}