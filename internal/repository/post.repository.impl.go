package repository

import (
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
func (p *PostRepositoryImpl) Create(req *requests.CreatePostReq, currentUser uint) (bool, error) {
	post := &models.Post{
		TextContent: req.TextContent,
		UserId:      currentUser,
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
			Id:        v.Id,
			TextContent:  v.TextContent,
			Author: responses.UserResponse{
				Id: v.UserId,
				UserName: v.User.UserName,
				FullName: v.User.FullName,
				Email: v.User.Email,
				AvatarUrl: v.User.AvatarUrl,
				Gender: v.User.Gender,
				CreatedAt: v.User.CreatedAt,
			},
			CreatedAt: v.CreatedAt,
		}
		result = append(result, post)
	}

	if res.Error != nil {
		return nil, res.Error
	}
	return result, nil
}

// Update implements PostRepository.
func (p *PostRepositoryImpl) Update(post *requests.CreatePostReq) (bool, error) {
	panic("unimplemented")
}
