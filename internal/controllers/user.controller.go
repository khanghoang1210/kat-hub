package controllers

import (
	"kathub/internal/models"
	"kathub/internal/services"
	"kathub/pkg/requests"
	"kathub/pkg/responses"

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

// GetAll		godoc
// @Summary			Get All Users
// @Description		Get All Users
// @Produce			application/json
// @Tags			users
// @Success      	200 {object} responses.ResponseData
// @Router			/users [get]
func (u UserController) GetAll(ctx *gin.Context) {
	result, _ := u.userService.GetAll()
	responses.APIResponse(ctx,result.StatusCode, result.Message, result.Data)
}

// Create		 	godoc
// @Summary			Create Users
// @Description		Create Users
// @Param			body body requests.CreateUserReq true "Users"
// @Produce			application/json
// @Tags			users
// @Success      	200 {object}  responses.ResponseData
// @Router			/users [post]
func (u UserController) Create(ctx *gin.Context) {
	body := requests.CreateUserReq{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		responses.APIResponse(ctx,400,"Bad request",nil)
		return
	}
	user := models.User{
		UserName: body.UserName,
		FullName: body.FullName,
		Email: body.Email,
		Password: body.Password,
	}
	result, _ := u.userService.Create(&user)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)

}
