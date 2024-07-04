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

