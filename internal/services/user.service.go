package services

import (
	"kathub/pkg/requests"
	response "kathub/pkg/responses"
)

type UserService interface {
	GetAll() (*response.ResponseData)
	Create(user *requests.CreateUserReq) (*response.ResponseData)
}