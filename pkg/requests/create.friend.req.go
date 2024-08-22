package requests

type CreateFriendReq struct {
	UserID int `json:"userID"`
	FriendID int `json:"friendID"`
}
