package services

import (
	"io"
	"kathub/pkg/requests"
	"kathub/pkg/responses"
)

type UserService interface {
	GetAll() (*responses.ResponseData)
	GetById(id uint)(*responses.ResponseData)
	Create(user *requests.CreateUserReq) (*responses.ResponseData)
	Update(user *requests.UpdateUserReq) (*responses.ResponseData)
	UploadAvatar(bucketName string, fileName string, data io.Reader) *responses.ResponseData
}
