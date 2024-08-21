package services

import "kathub/pkg/responses"

type FriendService interface {
	GetAll() *responses.ResponseData
	SendRequest(ReceiverID int) *responses.ResponseData
	RemoveRequest(id int) *responses.ResponseData
	AcceptRequest(id int) *responses.ResponseData
	RejectRequest(id int) *responses.ResponseData
}