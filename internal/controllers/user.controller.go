package controllers

import (
	"kathub/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}


func NewTagsController(service services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (u UserController)GetAll(ctx *gin.Context){
	userResponse := u.userService.GetAll()

	ctx.JSON(http.StatusOK, userResponse) 
}