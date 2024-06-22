package services

import (
	"kathub/internal/models"
	"kathub/internal/repository"

)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUsersServiceImpl(repository repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: repository,
	}
}

func (u UserServiceImpl)GetAll() []models.User{
	result:=u.userRepository.GetAll()
	var users []models.User

	users = append(users, result...)
	
	return users	
}