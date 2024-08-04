package repository

import (
	"kathub/internal/models"
	"kathub/pkg/requests"
	"kathub/pkg/responses"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetById(id uint) (*responses.UserResponse, error)
	Create(user *requests.CreateUserReq) (bool, error)
	Update(user *requests.UpdateUserReq) (bool, error)
	SaveAvatar(currentUser responses.UserResponse, avatarUrl string) (bool, error)
}
