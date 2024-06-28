package repository

import (
	"kathub/internal/models"
	"kathub/pkg/requests"
)

type AccountRepository interface {
	Login(loginReq *requests.LoginReq)(*models.User, error)
}