package services

import (
	"io"
	"kathub/internal/repository"
	"kathub/pkg/requests"
	"kathub/pkg/responses"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type PostServiceImpl struct {
	postRepo repository.PostRepository
	storage  StorageService
}

func NewPostServiceImpl(repo repository.PostRepository, storage StorageService) PostService {
	return &PostServiceImpl{
		postRepo: repo,
		storage:  storage,
	}
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
func (p *PostServiceImpl) Create(req *requests.CreatePostReq, currentUser responses.UserResponse) *responses.ResponseData {
	createdPostID, err := p.postRepo.Create(req, currentUser)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       createdPostID,
	}
}

func (p *PostServiceImpl) UploadPostImage(postID int, currentUser responses.UserResponse, data io.Reader, fileNameToUpload string) *responses.ResponseData {
	bucketName := "posts-image"
	extension := filepath.Ext(fileNameToUpload)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	fileName := "post" + "-" + strconv.Itoa(postID) + "-" + currentUser.UserName + "_" + timestamp + extension
	contentType := mime.TypeByExtension(extension)

	err := p.storage.Upload(bucketName, fileName, data, contentType)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	imageUrl := p.storage.GetPublicUrl(bucketName, fileName)

	result, insertImageErr := p.postRepo.InsertPostImage(postID, imageUrl)
	if insertImageErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    insertImageErr.Error(),
			Data:       false,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       result,
	}
}

// Update implements PostService.
func (p *PostServiceImpl) Update(req *requests.CreatePostReq, id uint) *responses.ResponseData {
	result, err := p.postRepo.Update(req, id)
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

func (p *PostServiceImpl) GetById(id uint) *responses.ResponseData {
	result, err := p.postRepo.GetById(id)

	if err != nil {
		if err.Error() == responses.StatusResourceNotFound {
			return &responses.ResponseData{
				StatusCode: http.StatusNoContent,
				Message:    err.Error(),
				Data:       false,
			}
		}
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

func (p *PostServiceImpl) Like(postID int, currentUser responses.UserResponse) *responses.ResponseData {
	result, err := p.postRepo.Like(postID, currentUser)

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

func (p *PostServiceImpl) UnLike(postID int, currentUser responses.UserResponse) *responses.ResponseData {
	result, err := p.postRepo.UnLike(postID, currentUser)

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

func (p *PostServiceImpl) CreateComment(req *requests.CreateCommentReq, currentUser responses.UserResponse) *responses.ResponseData {
	result, err := p.postRepo.InsertComment(req, currentUser)

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

func (p *PostServiceImpl) 	GetCommentsByPostID(postID int) *responses.ResponseData {
	result, err := p.postRepo.GetAllComment(postID)

	if err != nil {
		if err.Error() == responses.StatusResourceNotFound {
			return &responses.ResponseData{
				StatusCode: http.StatusNoContent,
				Message:    err.Error(),
				Data:       false,
			}
		}
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
