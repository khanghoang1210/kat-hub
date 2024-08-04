package repository

import (
	"errors"
	"fmt"
	"kathub/internal/models"
	"kathub/pkg/requests"
	"kathub/pkg/responses"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: Db}
}

func (u *UserRepositoryImpl) SaveAvatar(currentUser responses.UserResponse,avatarUrl string) (bool, error) {
	result := u.db.Model(&models.User{}).Where("id = ?", currentUser.Id).Update("avatar_url", avatarUrl)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
func (u *UserRepositoryImpl) GetById(id uint) (*responses.UserResponse, error) {
	var user models.User
	res := u.db.First(&user, id)
	if res.Error != nil {
		return nil, res.Error
	}
	result := &responses.UserResponse{
		Id:        user.Id,
		UserName:  user.UserName,
		FullName:  user.FullName,
		Email:     user.Email,
		Title:     user.Title,
		AvatarUrl: user.AvatarUrl,
		Gender:    user.Gender,
		CreatedAt: user.CreatedAt,
	}
	return result, nil
}

func (u *UserRepositoryImpl) GetAll() ([]models.User, error) {
	var users []models.User
	resp := u.db.Find(&users)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return users, nil
}

func (u *UserRepositoryImpl) Create(req *requests.CreateUserReq) (bool, error) {

	var users []models.User
	resp := u.db.Find(&users)
	if resp.Error != nil {
		return false, resp.Error
	}
	userFound := models.User{}
	existsRecord := u.db.Where("user_name = ?", req.UserName).First(&userFound)

	if existsRecord.RowsAffected != 0 {
		return false, fmt.Errorf("user name %s already exists", req.UserName)
	}

	user := &models.User{
		UserName: req.UserName,
		FullName: req.FullName,
		Password: req.Password,
		Email:    req.Email,
	}

	result := u.db.Create(user)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (u *UserRepositoryImpl) Update(req *requests.UpdateUserReq) (bool, error) {

	var user models.User
	resp := u.db.Where("user_name = ?", req.UserName).First(&user)
	if resp.Error != nil {

		if resp.RowsAffected == 0 {
			return false, errors.New(responses.StatusUserNotFound)
		}

		return false, resp.Error
	}

	user.FullName = req.FullName
	user.Email = req.Email
	user.Gender = req.Gender
	user.Title = req.Title

	result := u.db.Save(user)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
