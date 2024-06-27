package repository

import "kathub/pkg/requests"

type AccountRepository interface {
	Login(loginReq *requests.LoginReq)()
}