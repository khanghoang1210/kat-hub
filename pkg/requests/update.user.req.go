package requests

type UpdateUserReq struct {
	UserName string `json:"userName" binding:"required"`
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Title    string `json:"title" binding:"required"`
}
