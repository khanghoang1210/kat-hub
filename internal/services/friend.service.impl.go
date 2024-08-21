package services

import "kathub/pkg/responses"

type FriendServiceImpl struct {
}

func NewFriendRequestImpl() FriendService {
	return &FriendServiceImpl{}
}


// AcceptRequest implements FriendService.
func (f *FriendServiceImpl) AcceptRequest(id int) *responses.ResponseData {
	panic("unimplemented")
}

// GetAll implements FriendService.
func (f *FriendServiceImpl) GetAll() *responses.ResponseData {
	panic("unimplemented")
}

// RejectRequest implements FriendService.
func (f *FriendServiceImpl) RejectRequest(id int) *responses.ResponseData {
	panic("unimplemented")
}

// RemoveRequest implements FriendService.
func (f *FriendServiceImpl) RemoveRequest(id int) *responses.ResponseData {
	panic("unimplemented")
}

// SendRequest implements FriendService.
func (f *FriendServiceImpl) SendRequest(ReceiverID int) *responses.ResponseData {
	panic("unimplemented")
}

