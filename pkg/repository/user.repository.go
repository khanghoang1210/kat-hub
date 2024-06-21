package repository

import "kathub/internal/models"

type UserRepository interface {
	GetAll()[]models.User
}