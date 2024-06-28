package services

import (
	"kathub/pkg/requests"
	"kathub/pkg/responses"
)

type AccountService interface {
	Login(req *requests.LoginReq) (responses.ResponseData)
}