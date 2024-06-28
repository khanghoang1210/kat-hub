package repository

import (

	"fmt"
	"kathub/internal/models"
	"kathub/pkg/requests"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: Db}
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
	for _,v:= range users{
		if req.UserName == v.UserName{
			return false, fmt.Errorf("user name %s already exists", req.UserName)
		}
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
