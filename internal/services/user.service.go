package services

import (
	"kathub/pkg/requests"
	response "kathub/pkg/responses"
)

type UserService interface {
	GetAll() (*response.ResponseData)
	GetById(id uint)(*response.ResponseData)
	Create(user *requests.CreateUserReq) (*response.ResponseData)
	Update(user *requests.UpdateUserReq) (*response.ResponseData)
}
