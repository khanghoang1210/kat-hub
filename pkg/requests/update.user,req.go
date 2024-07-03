package requests

type UpdateUserReq struct {
	UserName string `json:"userName" binding:"required"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}
