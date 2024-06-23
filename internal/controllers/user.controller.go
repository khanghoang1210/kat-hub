package controllers

import (
	"kathub/internal/models"
	"kathub/internal/services"
	"kathub/pkg/response"
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

// @Router			/users [get]
func (u UserController)GetAll(ctx *gin.Context){
	userResponse, err := u.userService.GetAll()
	if err!=nil {
		response.ErrorResponse(ctx, 500, err.Error())
	}

	response.SuccessResponse(ctx,200, userResponse)
}

// Create		godoc
// @Summary			Create Users
// @Description		Create Users
// @Param			users body models.User true "Create Users"
// @Produce			application/json
// @Tags			users

// @Router			/users [post]
func (u UserController) Create(ctx *gin.Context){
	body:= models.User{}
	if err:=ctx.BindJSON(&body);err!=nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
		return
	 }
	//user := models.User{}
	result, err := u.userService.Create(&body)
	if err!=nil{
		response.ErrorResponse(ctx, 500, err.Error())
	}
	response.SuccessResponse(ctx,200,result)

}