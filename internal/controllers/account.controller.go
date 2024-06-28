package controllers

import (
	"kathub/internal/services"
	"kathub/pkg/requests"
	"kathub/pkg/responses"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService services.AccountService
}

func NewAccountController(service services.AccountService) *AccountController {
	return &AccountController{
		accountService: service,
	}
}

// GetAll		godoc
// @Summary			Login
// @Description		Login
// @Param			body body requests.LoginReq true "Accounts"
// @Produce			application/json
// @Tags			accounts
// @Success      	200 {object} responses.ResponseData
// @Router			/accounts/login [post]
func (as AccountController) Login(ctx *gin.Context) {
	loginReq := requests.LoginReq{}
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}
	result := as.accountService.Login(&loginReq)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}
