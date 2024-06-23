package services

import "kathub/internal/models"

type UserService interface {
	GetAll() ([]models.User, error)
	Create(user *models.User) (bool, error)
}