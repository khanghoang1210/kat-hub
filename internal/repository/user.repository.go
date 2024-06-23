package repository

import "kathub/internal/models"

type UserRepository interface {
	GetAll()([]models.User, error)
	Create(user *models.User) (bool,error)
}