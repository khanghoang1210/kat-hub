package services

import (
	"kathub/internal/repository"
	"kathub/pkg/requests"
	response "kathub/pkg/responses"

	"golang.org/x/crypto/bcrypt"
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
func (u UserServiceImpl)Create(user *requests.CreateUserReq)  (*response.ResponseData, error){
	response := &response.ResponseData{}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password),10) 
	if err!=nil{
		response.StatusCode = 500
		response.Message = "Fail to hash password"
		response.Data = false
		return response, err
	}
	user.Password = string(passwordHash)

	result, err := u.userRepository.Create(user)


	if err!=nil {
		response.StatusCode = 500
		response.Message = "Internal Error"
		response.Data = false
		return response, err
	}
	response.StatusCode = 200
	response.Message = "Success"
	response.Data = result
	return response, nil
	
}