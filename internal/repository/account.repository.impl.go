package repository

import (
	"kathub/pkg/requests"
	"kathub/pkg/responses"

	"gorm.io/gorm"
)

type AccountRepositoryImpl struct {
	db *gorm.DB
}

func AccountUserRepositoryImpl(Db *gorm.DB) AccountRepository {
	return &AccountRepositoryImpl{db: Db}
}

func (ar *AccountRepositoryImpl)Login(loginReq *requests.LoginReq)(*responses.LoginRes, error){
	return nil,nil
}