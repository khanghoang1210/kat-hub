package repository

import (
	"kathub/internal/models"
	"kathub/pkg/requests"
	"kathub/pkg/responses"

	"gorm.io/gorm"
)

type FriendRepositoryImpl struct {
	db *gorm.DB
}

func NewFriendRepositoryImpl(db *gorm.DB) FriendRepository {
	return &FriendRepositoryImpl{
		db: db,
	}
}

// Create implements FriendRepository.
func (f *FriendRepositoryImpl) Create(req requests.CreateFriendReq) (bool, error) {
	relationRight := &models.Friend{
		UserId:   uint(req.UserID),
		FriendId: uint(req.FriendID),
	}
	relationLeft := &models.Friend{
		UserId:   uint(req.FriendID),
		FriendId: relationRight.UserId,
	}
	result1 := f.db.Create(&relationRight)
	if result1.Error != nil {
		return false, result1.Error
	}
	result2 := f.db.Create(&relationLeft)
	if result2.Error != nil {
		return false, result2.Error
	}
	return true, nil
}

// CreateRequest implements FriendRepository.
func (f *FriendRepositoryImpl) CreateRequest(req requests.CreateRequestReq, currentUser int) (bool, error) {
	request := &models.Request{
		SenderID:   uint(currentUser),
		ReceiverID: uint(req.ReceiverID),
		Status:     1,
	}
	result := f.db.Create(request)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (f *FriendRepositoryImpl) AcceptRequest(id int) (bool, error) {
	request := &models.Request{}

	res := f.db.First(&request, id)
	if res.Error != nil {
		return false, res.Error
	}
	request.Status = 2
	result := f.db.Save(&request)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil

}

// DeclineRequest implements FriendRepository.
func (f *FriendRepositoryImpl) DeclineRequest(id int) (bool, error) {
	request := &models.Request{}

	res := f.db.First(&request, id)
	if res.Error != nil {
		return false, res.Error
	}
	request.Status = 3
	result := f.db.Save(&request)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// Delete implements FriendRepository.
func (f *FriendRepositoryImpl) Delete(friendID int, currentUser int) (bool, error) {
	relationLeft := &models.Friend{}
	relationRight := &models.Friend{}

	f.db.Where("user_id", currentUser).Where("friend_id", friendID).Find(&relationLeft)
	f.db.Where("user_id", currentUser).Where("friend_id", friendID).Find(&relationRight)

	result1 := f.db.Delete(&relationLeft)
	result2 := f.db.Delete(&relationRight)

	if result1.Error != nil {
		return false, result1.Error
	}
	if result2.Error != nil {
		return false, result2.Error
	}
	return true, nil
}

// DeleteRequest implements FriendRepository.
func (f *FriendRepositoryImpl) DeleteRequest(id int) (bool, error) {
	request := &models.Request{}

	res := f.db.First(&request, id)
	if res.Error != nil {
		return false, res.Error
	}
	result := f.db.Delete(request)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil

}

// GetAll implements FriendRepository.
func (f *FriendRepositoryImpl) GetAll(userID int) ([]responses.UserResponse, error) {
	friends := []models.Friend{}
	resp := f.db.Find(&friends, userID)
	if resp.Error != nil {
		return nil, resp.Error
	}

	result := []responses.UserResponse{}
	for _, v := range friends {
		friend := responses.UserResponse{
			Id:        v.FriendId,
			UserName:  v.Friend.UserName,
			FullName:  v.Friend.FullName,
			Email:     v.Friend.Email,
			Title:     v.Friend.Title,
			AvatarUrl: v.Friend.AvatarUrl,
			CreatedAt: v.Friend.CreatedAt,
		}

		result = append(result, friend)
	}
	return result, nil
}

func (f *FriendRepositoryImpl) GetAllRequests(userID int) ([]responses.RequestRes, error) {
	requests := []models.Request{}
	resp := f.db.Find(&requests, userID)
	if resp.Error != nil {
		return nil, resp.Error
	}
	result := []responses.RequestRes{}
	for _, v := range requests {
		request := &responses.RequestRes{
			Id:         int(v.Id),
			SenderID:   int(v.SenderID),
			ReceiverID: int(v.ReceiverID),
			CreatedAt:  v.CreatedAt,
		}

		result = append(result, *request)
	}
	return result, nil
}

func (f *FriendRepositoryImpl) GetRequestByID(id int) (*responses.RequestRes, error) {
	request := &models.Request{}
	resp := f.db.First(&request, id)
	if resp.Error != nil {
		return nil, resp.Error
	}

	result := &responses.RequestRes{
		Id:         int(request.Id),
		SenderID:   int(request.SenderID),
		ReceiverID: int(request.ReceiverID),
		Status:     request.Status,
		CreatedAt:  request.CreatedAt,
	}
	return result, nil
}
