package controllers

import (
	"kathub/internal/services"
	"kathub/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}


func NewUserController(service services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (u UserController)GetAll(ctx *gin.Context){
	userResponse := u.userService.GetAll()

	response.SuccessResponse(ctx,200, userResponse)
}