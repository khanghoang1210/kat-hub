package repository

import (
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

func(u *UserRepositoryImpl) GetAll() ([]models.User, error){
	var users []models.User
	resp :=u.db.Find(&users)
	if resp.Error != nil{
		return nil, resp.Error
	}

	return users,nil
}

func(u *UserRepositoryImpl)Create(user *requests.CreateUserReq) (bool, error){


	result := u.db.Create(user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}