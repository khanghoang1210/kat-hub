package services

import (
	"kathub/pkg/requests"
	response "kathub/pkg/responses"
)

type UserService interface {
	GetAll() (*response.ResponseData, error)
	Create(user *requests.CreateUserReq) (*response.ResponseData, error)
}