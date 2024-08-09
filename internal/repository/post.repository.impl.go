package repository

import (
	"errors"
	"kathub/internal/models"
	"kathub/pkg/requests"
	"kathub/pkg/responses"

	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepositoryImpl(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{db: db}
}

// Create implements PostRepository.
func (p *PostRepositoryImpl) Create(req *requests.CreatePostReq, currentUser responses.UserResponse) (bool, error) {
	post := &models.Post{
		TextContent: req.TextContent,
		UserId:      currentUser.Id,
	}
	result := p.db.Create(post)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// GetAll implements PostRepository.
func (p *PostRepositoryImpl) GetAll() ([]*responses.PostResponse, error) {
	posts := []*models.Post{}

	res := p.db.Preload("User").Order("created_at DESC").Find(&posts)
	var result []*responses.PostResponse

	for _, v := range posts {

		post := &responses.PostResponse{
			Id:          v.Id,
			TextContent: v.TextContent,
			Author: responses.UserResponse{
				Id:        v.UserId,
				UserName:  v.User.UserName,
				FullName:  v.User.FullName,
				Email:     v.User.Email,
				Title:     v.User.Title,
				AvatarUrl: v.User.AvatarUrl,
				Gender:    v.User.Gender,
				CreatedAt: v.User.CreatedAt,
			},
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		result = append(result, post)
	}

	if res.Error != nil {
		return nil, res.Error
	}
	return result, nil
}

// Update implements PostRepository.
func (p *PostRepositoryImpl) Update(req *requests.CreatePostReq, id uint) (bool, error) {
	post := models.Post{}
	res := p.db.First(&post, id)
	if res.Error != nil {

		if res.RowsAffected == 0 {
			return false, errors.New(responses.StatusResourceNotFound)
		}

		return false, res.Error
	}
	post.TextContent = req.TextContent
	result := p.db.Save(&post)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (p *PostRepositoryImpl) Delete(id uint) (bool, error) {
	res := p.db.Delete(&models.Post{}, id)

	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (p *PostRepositoryImpl) GetById(id uint) (*responses.PostResponse, error) {
	var post models.Post
	res := p.db.Preload("User").First(&post, id)
	if res.Error != nil {

		if res.RowsAffected == 0 {
			return nil, errors.New(responses.StatusResourceNotFound)
		}

		return nil, res.Error
	}
	result := &responses.PostResponse{
		Id:          post.Id,
		TextContent: post.TextContent,
		Author: responses.UserResponse{
			Id:        post.UserId,
			UserName:  post.User.UserName,
			FullName:  post.User.FullName,
			Email:     post.User.Email,
			Title:     post.User.Title,
			AvatarUrl: post.User.AvatarUrl,
			Gender:    post.User.Gender,
			CreatedAt: post.User.CreatedAt,
		},
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return result, nil
}

func (p *PostRepositoryImpl) InsertPostImage(postID uint, imageUrl string) (bool, error) {
	result := p.db.Model(&models.Post{}).Where("id = ?", postID).Update("imageContent", imageUrl)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (p *PostRepositoryImpl) Like(id uint, user responses.UserResponse) (bool, error) {
	return true, nil
}

func (p *PostRepositoryImpl) UnLike(id uint, user responses.UserResponse) (bool, error) {
	return true, nil
}
