package services

import (
	"kathub/internal/repository"
	"kathub/pkg/requests"
	"kathub/pkg/responses"
	"net/http"
)

type FriendServiceImpl struct {
	repo repository.FriendRepository
}

func NewFriendRequestImpl(repo repository.FriendRepository) FriendService {
	return &FriendServiceImpl{
		repo: repo,
	}
}


// AcceptRequest implements FriendService.
func (f *FriendServiceImpl) AcceptRequest(id int) *responses.ResponseData {
	resultRequest, errRequest := f.repo.AcceptRequest(id) 
	if errRequest != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    errRequest.Error(),
			Data:       nil,
		}
	}
	req,  err := f.repo.GetRequestByID(id)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusNotFound,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	obj := &requests.CreateFriendReq{
		UserID: req.ReceiverID,
		FriendID: req.SenderID,
	}

	resultFriend, errFriend := f.repo.Create(*obj)

	if errFriend != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    errFriend.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       resultFriend && resultRequest,
	}
}

// GetAll implements FriendService.
func (f *FriendServiceImpl) GetAll(currentUser int) *responses.ResponseData {
	result, err := f.repo.GetAll(currentUser)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	if len(result) == 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusNoContent,
			Message:    responses.StatusNoData,
			Data:       nil,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       result,
	}
}

// RejectRequest implements FriendService.
func (f *FriendServiceImpl) RejectRequest(id int) *responses.ResponseData {
	result, err := f.repo.DeclineRequest(id)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       result,
	}
}

// RemoveRequest implements FriendService.
func (f *FriendServiceImpl) RemoveRequest(id int) *responses.ResponseData {
	result, err := f.repo.DeleteRequest(id)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       result,
	}
}

// SendRequest implements FriendService.
func (f *FriendServiceImpl) SendRequest(receiverID int, currentUser int) *responses.ResponseData {
	req := &requests.CreateRequestReq{
		ReceiverID: receiverID,
	}
	result, err := f.repo.CreateRequest(*req, currentUser)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       result,
	}
}

