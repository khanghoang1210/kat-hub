package repository

import (
	"kathub/internal/models"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: Db}
}

func(u *UserRepositoryImpl) GetAll() []models.User{
	var users []models.User
	resp :=u.db.Find(&users)
	if resp.Error != nil{
		return nil
	}

	return users
}