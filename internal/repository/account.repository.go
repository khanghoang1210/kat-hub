package repository

import (
	"kathub/pkg/requests"
	"kathub/pkg/responses"
)

type AccountRepository interface {
	Login(loginReq *requests.LoginReq)(*responses.LoginRes, error)
}