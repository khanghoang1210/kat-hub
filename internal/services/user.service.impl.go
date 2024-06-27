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

func (u UserServiceImpl)GetAll()  (*response.ResponseData, error){
	result, err:=u.userRepository.GetAll()
	if err!=nil {
		return nil, err
	}
	var users []*response.UserResponse

	for _,v:=range result{
		user := &response.UserResponse{
			Id: v.Id,
			FullName: v.FullName,
			UserName: v.UserName,
			Email: v.Email,
			AvatarUrl: v.AvatarUrl,
			CreatedAt: v.CreatedAt,
		}
		users = append(users, user)
	}
	response := &response.ResponseData{
		StatusCode: 200,
		Message: "Success",
		Data: users,
	}

	return response, nil
}
func (u UserServiceImpl)Create(user *models.User)  (*response.ResponseData, error){
	result, err := u.userRepository.Create(user)
	response := &response.ResponseData{}
	if err!=nil {
		response.StatusCode = 500
		response.Message = "Fail"
		response.Data = false
		return response, err
	}
	response.StatusCode = 200
	response.Message = "Success"
	response.Data = result
	return response, nil
	
}