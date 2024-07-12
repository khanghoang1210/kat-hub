package repository

import (
	"kathub/internal/models"
	"kathub/pkg/requests"

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
func (p *PostRepositoryImpl) GetAll() ([]*models.Post, error) {
	posts := []*models.Post{}

	res := p.db.Find(&posts)

	if res.Error != nil {
		return nil, res.Error
	}
	return posts, nil
}

// Update implements PostRepository.
func (p *PostRepositoryImpl) Update(post *requests.CreatePostReq) (bool, error) {
	panic("unimplemented")
}
