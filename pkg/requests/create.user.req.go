package requests

type CreateUserReq struct {
	UserName        string `json:"userName" validate:"required"`
	FullName        string `json:"fullName" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}
