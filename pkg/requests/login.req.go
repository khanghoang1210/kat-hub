package requests

type LoginReq struct {
	UserName        string `json:"userName" binding:"required"`
	Password        string `json:"password" binding:"required"`
}
