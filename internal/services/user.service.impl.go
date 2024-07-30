package services

import (
	"io"
	"kathub/internal/repository"
	"kathub/pkg/requests"
	"kathub/pkg/responses"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	storage             StorageService
}

func NewUsersServiceImpl(repository repository.UserRepository, storage StorageService) UserService {
	return &UserServiceImpl{
		userRepository: repository,
		storage:             storage,
	}
}

// GetById implements UserService.
func (u *UserServiceImpl) GetById(id uint) *responses.ResponseData {
	result, err := u.userRepository.GetById(id)

	if err != nil {
		if err.Error() == responses.StatusResourceNotFound {
			return &responses.ResponseData{
				StatusCode: http.StatusNoContent,
				Message:    err.Error(),
				Data:       false,
			}
		}
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       result,
	}
}

func (u UserServiceImpl) GetAll() *responses.ResponseData {
	result, err := u.userRepository.GetAll()
	if err != nil {
		return &responses.ResponseData{
			StatusCode: 500,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	var users []responses.UserResponse

	for _, v := range result {
		user := responses.UserResponse{
			Id:        v.Id,
			FullName:  v.FullName,
			UserName:  v.UserName,
			Email:     v.Email,
			Gender:    v.Gender,
			AvatarUrl: v.AvatarUrl,
			CreatedAt: v.CreatedAt,
		}
		users = append(users, user)
	}
	return &responses.ResponseData{
		StatusCode: 200,
		Message:    "Success",
		Data:       users,
	}
}

func (u UserServiceImpl) Create(user *requests.CreateUserReq) *responses.ResponseData {

	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errHash != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Fail to hash password",
			Data:       false,
		}
	}

	user.Password = string(passwordHash)

	result, err := u.userRepository.Create(user)

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       result,
	}

}

func (us UserServiceImpl) Update(user *requests.UpdateUserReq) *responses.ResponseData {

	userUpdated, err := us.userRepository.Update(user)
	if err != nil {
		if err.Error() == responses.StatusUserNotFound {
			return &responses.ResponseData{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       false,
			}
		}
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       userUpdated,
	}
}

func (us UserServiceImpl) UploadAvatar(fileName io.Reader) *responses.ResponseData {
	bucketName := "test"

	err := us.storage.Upload(bucketName, fileName)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       true,
	}
}
