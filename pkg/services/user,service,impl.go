package services

import (
	"kathub/internal/models"
	"kathub/pkg/repository"

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
	for _,value := range result{
		users = append(users, value)
	}
	return users	
}