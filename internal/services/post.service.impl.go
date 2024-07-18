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

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       result,
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
// Update implements PostService.
func (p *PostServiceImpl) Update(req *requests.CreatePostReq, id uint) *responses.ResponseData {
	result, err := p.postRepo.Update(req,id)
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

func (p *PostServiceImpl) Delete(id uint) *responses.ResponseData {
	result, err := p.postRepo.Delete(id)

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       result,
	}

}
