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

func (u UserServiceImpl)GetAll() ([]models.User, error){
	result, err:=u.userRepository.GetAll()
	if err!=nil {
		return nil, err
	}
	var users []models.User

	users = append(users, result...)
	
	return users, nil	
}
func (u UserServiceImpl)Create(user *models.User) (bool, error){
	result, err := u.userRepository.Create(user)
	if err!=nil {
		return false, err
	}
	return result, nil
	
}