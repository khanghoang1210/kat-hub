package controllers

import (
	"kathub/internal/models"
	"kathub/internal/services"
	"kathub/pkg/requests"
	"kathub/pkg/responses"
	"net/http"

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
// @Success      	200  {object} response.ResponseData{data=models.User}
// @Router			/users [get]
func (u UserController) GetAll(ctx *gin.Context) {
	userResponse := u.userService.GetAll()
	ctx.JSON(http.StatusOK, userResponse)
}

// Create		 	godoc
// @Summary			Create Users
// @Description		Create Users
// @Param			body body requests.CreateUserReq true "Users"
// @Produce			application/json
// @Tags			users
// @Success      	200  {object}  response.ResponseData{data=bool}
// @Router			/users [post]
func (u UserController) Create(ctx *gin.Context) {
	body := requests.CreateUserReq{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseData{
			StatusCode: 400,
			Message: err.Error(),
			Data: false,
		})
		return
	}
	user := models.User{
		UserName: body.UserName,
		FullName: body.FullName,
		Email: body.Email,
		Password: body.Password,
	}
	result := u.userService.Create(&user)
	ctx.JSON(http.StatusOK, result)

}
