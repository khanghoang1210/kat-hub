package controllers

import (
	"fmt"
	"kathub/internal/services"
	"kathub/pkg/requests"
	"kathub/pkg/responses"
	"kathub/pkg/utils"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(service services.PostService) *PostController {
	return &PostController{
		postService: service,
	}
}

// Create		 	godoc
// @Summary			Create Posts
// @Description		Create Posts
// @Param			body body requests.CreatePostReq true "Posts"
// @Produce			application/json
// @Tags			Posts
// @Success      	200 {object}  responses.ResponseData
// @Router			/posts [post]
func (pc PostController) Create(ctx *gin.Context) {
	newPost := requests.CreatePostReq{}
	currentUser := utils.GetUserProfile(ctx)
	fmt.Print(currentUser)
	if currentUser != nil {
		fmt.Print(currentUser)
		return
	}

	if err := ctx.ShouldBindJSON(&newPost); err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	result := pc.postService.Create(&newPost, 1)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)

}
