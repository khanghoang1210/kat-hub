package controllers

import (
	"kathub/internal/services"
	"kathub/pkg/requests"
	"kathub/pkg/responses"
	"kathub/pkg/utils"
	"strconv"

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
// @Param			body body requests.CreatePostReq true "Posts"
// @Produce			application/json
// @Tags			Posts
// @Success      	200 {object}  responses.ResponseData
// @Router			/posts [post]
func (pc PostController) Create(ctx *gin.Context) {
	newPost := requests.CreatePostReq{}
	currentUser, err := utils.GetUserProfile(ctx)

	if err != nil {
		responses.APIResponse(ctx, 401, responses.StatusUnAuthorize, nil)
	}

	if err := ctx.ShouldBindJSON(&newPost); err != nil {
		responses.APIResponse(ctx, 400, responses.StatusParamInvalid, nil)
		return
	}

	result := pc.postService.Create(&newPost, *currentUser)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)

}

func (pc PostController) UploadPostImage(ctx *gin.Context) {
	file, fileErr := ctx.FormFile("imageContent")
	if fileErr != nil {
		responses.APIResponse(ctx, 400, responses.StatusParamInvalid, nil)
		return
	}

	id, errParse := strconv.Atoi(ctx.Param("id"))
	if errParse != nil {
		responses.APIResponse(ctx, 400, responses.StatusParamInvalid, nil)
		return
	}
	currentUser, err := utils.GetUserProfile(ctx)

	if err != nil {
		responses.APIResponse(ctx, 401, responses.StatusUnAuthorize, nil)
	}

	f, errFile := file.Open()
	if errFile != nil {
		responses.APIResponse(ctx, 500, responses.StatusInternalError, nil)
	}
	defer f.Close()

	result := pc.postService.UploadPostImage(id, *currentUser, f, file.Filename)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

// GetAll		 	godoc
// @Summary			Get All Posts
// @Produce			application/json
// @Tags			Posts
// @Success      	200 {object}  responses.ResponseData
// @Router			/posts [get]
func (pc PostController) GetAll(ctx *gin.Context) {
	result := pc.postService.GetAll()
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

// GetAll		 	godoc
// @Summary			Update Post
// @Produce			application/json
// @Param 			id  path  int  true  "Post ID"
// @Param			body body requests.CreatePostReq true "Posts"
// @Tags			Posts
// @Success      	200 {object}  responses.ResponseData
// @Router			/posts [put]
func (pc PostController) Update(ctx *gin.Context) {
	updatePost := requests.CreatePostReq{}
	id, errParse := strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&updatePost); errParse != nil || err != nil {
		responses.APIResponse(ctx, 400, responses.StatusParamInvalid, nil)
		return
	}
	result := pc.postService.Update(&updatePost, uint(id))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

// GetAll		 	godoc
// @Summary			Delete Post
// @Produce			application/json
// @Param 			id  path  int  true  "Post ID"
// @Tags			Posts
// @Success      	200 {object}  responses.ResponseData
// @Router			/posts [delete]
func (pc PostController) Delete(ctx *gin.Context) {
	id, errParse := strconv.Atoi(ctx.Param("id"))

	if errParse != nil {
		responses.APIResponse(ctx, 400, responses.StatusParamInvalid, nil)
		return
	}
	result := pc.postService.Delete(uint(id))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

// GetAll		 	godoc
// @Summary			Get Post By Id
// @Produce			application/json
// @Param 			id  path  int  true  "Post ID"
// @Tags			Posts
// @Success      	200 {object}  responses.ResponseData
// @Router			/posts/{id} [get]
func (pc PostController) GetById(ctx *gin.Context) {

	id, errParse := strconv.Atoi(ctx.Param("id"))
	if errParse != nil {
		responses.APIResponse(ctx, 400, responses.StatusParamInvalid, nil)
		return
	}

	result := pc.postService.GetById(uint(id))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (pc PostController) Like(ctx *gin.Context) {
	currentUser, err := utils.GetUserProfile(ctx)
	postID, errParse := strconv.Atoi(ctx.Param("id"))
	if errParse != nil || err != nil {
		responses.APIResponse(ctx, 400, responses.StatusParamInvalid, nil)
		return
	}
	result := pc.postService.Like(postID, *currentUser)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (pc PostController) UnLike(ctx *gin.Context) {
	currentUser, err := utils.GetUserProfile(ctx)
	postID, errParse := strconv.Atoi(ctx.Param("id"))
	if errParse != nil || err != nil {
		responses.APIResponse(ctx, 400, responses.StatusParamInvalid, nil)
		return
	}
	result := pc.postService.UnLike(postID, *currentUser)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func(pc PostController) CreateComment(ctx *gin.Context) {
	newComment := requests.CreateCommentReq{}
	currentUser,_ := utils.GetUserProfile(ctx)

	if err := ctx.ShouldBindJSON(&newComment); err != nil {
		responses.APIResponse(ctx, 400, responses.StatusParamInvalid, nil)
		return
	}

	result := pc.postService.CreateComment(&newComment, *currentUser)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (pc PostController) GetCommentsByPostID(ctx *gin.Context){
	postID, errParse := strconv.Atoi(ctx.Param("id"))
	if errParse != nil {
		responses.APIResponse(ctx, 400, responses.StatusParamInvalid, nil)
		return
	}

	result := pc.postService.GetCommentsByPostID(postID)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (pc PostController) EditComment(ctx *gin.Context){
	updateComment := requests.UpdateCommentReq{}
	currentUser,_ := utils.GetUserProfile(ctx)

	if err := ctx.ShouldBindJSON(&updateComment); err != nil {
		responses.APIResponse(ctx, 400, responses.StatusParamInvalid, nil)
		return
	}
	result := pc.postService.UpdateComment(&updateComment, *currentUser)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (pc PostController) DeleteComment(ctx *gin.Context){

}
