package repository

import (
	"kathub/pkg/requests"
	"kathub/pkg/responses"
)

type FriendRepository interface {
	GetAll(userID int) ([]responses.UserResponse, error)
	Create(req requests.CreateFriendReq) (bool, error)
	Delete(friendID int, currentUser int) (bool, error)
	CreateRequest(req requests.CreateRequestReq, currentUser int) (bool, error)
	DeleteRequest(id int) (bool, error)
	AcceptRequest(id int) (bool, error)
	DeclineRequest(id int) (bool, error)
}