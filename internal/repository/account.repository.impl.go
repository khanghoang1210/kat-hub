package repository

import (
	"errors"
	"kathub/internal/models"
	"kathub/pkg/requests"
	"kathub/pkg/responses"

	"gorm.io/gorm"
)

type AccountRepositoryImpl struct {
	db *gorm.DB
}

func NewAccountRepositoryImpl(Db *gorm.DB) AccountRepository {
	return &AccountRepositoryImpl{db: Db}
}

func (ar *AccountRepositoryImpl) Login(loginReq *requests.LoginReq) (*models.User, error) {
	user := &models.User{}
	result := ar.db.Where("user_name = ?", loginReq.UserName).First(&user)
	
	if result.RowsAffected == 0 {
		return nil, errors.New(responses.StatusUserNotFound)
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil

}
