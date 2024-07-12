package services

import (
	"kathub/internal/repository"
	"kathub/pkg/requests"
	"kathub/pkg/responses"
	"net/http"

)

type PostServiceImpl struct {
	postRepo repository.PostRepository
}



func NewPostServiceImpl(repo repository.PostRepository) PostService {
	return &PostServiceImpl{postRepo: repo}
}

// GetAll implements PostService.
func (p *PostServiceImpl) GetAll() *responses.ResponseData {
	result, err := p.postRepo.GetAll()

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	if len(result) == 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusNoContent,
			Message:    responses.StatusNoData,
			Data:       nil,
		}
	}
	var posts []responses.PostResponse

	for _, v := range result {
		post := responses.PostResponse{
			Id:        v.Id,
			TextContent:  v.TextContent,
			AuthorID: v.UserId,
			CreatedAt: v.CreatedAt,
		}
		posts = append(posts, post)
	}
	return  &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       posts,
	}
}

// Create implements PostService.
func (p *PostServiceImpl) Create(req *requests.CreatePostReq, currentUser uint) *responses.ResponseData {

	result, err := p.postRepo.Create(req, currentUser)

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       result,
	}
}
