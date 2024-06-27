package services

import (
	"kathub/internal/models"
	response "kathub/pkg/responses"
)

type UserService interface {
	GetAll() (*response.ResponseData, error)
	Create(user *models.User) (*response.ResponseData, error)
}