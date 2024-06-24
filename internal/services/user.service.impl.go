package services

import (
	"kathub/internal/models"
	"kathub/internal/repository"
	response "kathub/pkg/responses"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUsersServiceImpl(repository repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: repository,
	}
}

func (u UserServiceImpl)GetAll() *response.ResponseData{
	result, err:=u.userRepository.GetAll()
	if err!=nil {
		return &response.ResponseData{
			StatusCode: 500,
			Message: err.Error(),
			Data: nil,
		}
	}
	var users []models.User

	users = append(users, result...)
	
	return &response.ResponseData{
		StatusCode: 200,
		Message: "Success",
		Data: users,
	}
}
func (u UserServiceImpl)Create(user *models.User) *response.ResponseData{
	result, err := u.userRepository.Create(user)
	if err!=nil {
		return &response.ResponseData{
			StatusCode: 500,
			Message: err.Error(),
			Data: nil,
		}
	}
	return &response.ResponseData{
		StatusCode: 200,
		Message: "Success",
		Data: result,
	}
	
}