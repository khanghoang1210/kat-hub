package repository

import "kathub/pkg/responses"

type FriendRepository interface {
	GetAll(userID int) ([]responses.UserResponse, error)
	Create()
	Delete()
	CreateRequest(receiverID int) (bool, error)
	DeleteRequest(id int) (bool, error)
}