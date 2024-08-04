package controllers

import (
	"kathub/internal/services"
	"kathub/pkg/requests"
	"kathub/pkg/responses"
	"kathub/pkg/utils"
	"strconv"

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
// @Produce			application/json
// @Tags			users
// @Success      	200 {object} responses.ResponseData
// @Router			/users [get]
func (u UserController) GetAll(ctx *gin.Context) {
	result := u.userService.GetAll()
	responses.APIResponse(ctx,result.StatusCode, result.Message, result.Data)
}

// Create		 	godoc
// @Summary			Create Users
// @Param			body body requests.CreateUserReq true "Users"
// @Produce			application/json
// @Tags			users
// @Success      	200 {object}  responses.ResponseData
// @Router			/users [post]
func (u UserController) Create(ctx *gin.Context) {
	newUser := requests.CreateUserReq{}
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		responses.APIResponse(ctx,400,"Bad request",nil)
		return
	}
	
	result := u.userService.Create(&newUser)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)

}

// GetAll		godoc
// @Summary			Update Users
// @Param			body body requests.UpdateUserReq true "Users"
// @Produce			application/json
// @Tags			users
// @Success      	200 {object} responses.ResponseData
// @Router			/users [put]
func (u UserController) Update(ctx *gin.Context) {
	updatedUser := requests.UpdateUserReq{}
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		responses.APIResponse(ctx,400,"Bad request",nil)
		return
	}
	
	result := u.userService.Update(&updatedUser)
	responses.APIResponse(ctx,result.StatusCode, result.Message, result.Data)
}

// GetAll		godoc
// @Summary			Get User By ID
// @Param			id  path  int  true  "User ID"
// @Produce			application/json
// @Tags			users
// @Success      	200 {object} responses.ResponseData
// @Router			/users/{id} [get]
func (u UserController) GetById(ctx *gin.Context){
	id, errParse := strconv.Atoi(ctx.Param("id"))

	if  errParse != nil {
		responses.APIResponse(ctx, 400, responses.StatusParamInvalid, nil)
		return
	}
	result := u.userService.GetById(uint(id))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (u UserController) UploadAvatar(ctx * gin.Context){
	file, err := ctx.FormFile("avatar")
	currentUser, getUserErr := utils.GetUserProfile(ctx)
	if getUserErr != nil {
		responses.APIResponse(ctx, 403, responses.StatusUnAuthorize, nil)
	}
	if err != nil {
		responses.APIResponse(ctx, 400, responses.StatusParamInvalid, nil)
	 return
	}

	f, errFile := file.Open()
	if errFile != nil {
		responses.APIResponse(ctx, 500, responses.StatusInternalError, nil)
	}
	defer f.Close()
	result := u.userService.UploadAvatar(*currentUser,f, file.Filename)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)

}