package repository

import (
	"kathub/internal/models"
	"kathub/pkg/requests"
)

type UserRepository interface {
	GetAll()([]models.User, error)
	Create(user *requests.CreateUserReq) (bool,error)
	Update(user *requests.UpdateUserReq) (bool, error)
}