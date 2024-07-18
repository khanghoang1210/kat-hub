package services

import (
	"kathub/pkg/requests"
	"kathub/pkg/responses"
)

type PostService interface {
	Create(req *requests.CreatePostReq, currentUser uint) *responses.ResponseData
	GetAll() *responses.ResponseData
	Update(req *requests.CreatePostReq, id uint) *responses.ResponseData
	Delete(id uint) *responses.ResponseData
}