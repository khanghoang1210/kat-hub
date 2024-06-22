package services

import "kathub/internal/models"

type UserService interface {
	GetAll() []models.User
}