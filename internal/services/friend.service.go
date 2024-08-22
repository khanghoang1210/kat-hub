package services

import "kathub/pkg/responses"

type FriendService interface {
	GetAll(currentUser int) *responses.ResponseData
	SendRequest(receiverID int, currentUser int) *responses.ResponseData
	RemoveRequest(id int) *responses.ResponseData
	AcceptRequest(id int) *responses.ResponseData
	RejectRequest(id int) *responses.ResponseData
}