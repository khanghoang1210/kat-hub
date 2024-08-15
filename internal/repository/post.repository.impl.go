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
func (p *PostRepositoryImpl) Create(req *requests.CreatePostReq, currentUser responses.UserResponse) (*uint, error) {
	post := &models.Post{
		TextContent: req.TextContent,
		UserId:      currentUser.Id,
	}
	result := p.db.Create(post)

	if result.Error != nil {
		return nil, result.Error
	}
	return &post.Id, nil
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
			ImageContent: v.ImageContent,
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

func (p *PostRepositoryImpl) InsertPostImage(postID int, imageUrl string) (bool, error) {
	var post models.Post
	if err := p.db.Where("id = ?", postID).First(&post).Error; err != nil {
		return false, err
	}

	//post.ImageContent = append(post.ImageContent, imageUrl)
	//imageContentArray := pq.Array(post.ImageContent)

	result := p.db.Model(&models.Post{}).Where("id = ?", postID).Update("image_content", post.ImageContent)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (p *PostRepositoryImpl) Like(postID int, user responses.UserResponse) (bool, error) {
	like := &models.Like{
		UserId: user.Id,
		PostId: uint(postID),
	}

	result := p.db.Create(like)

	if result.Error != nil {
		return true, result.Error
	}
	
	return true, nil
}

func (p *PostRepositoryImpl) UnLike(postID int, user responses.UserResponse) (bool, error) {
	res := p.db.Where("user_id", user.Id).Where("post_id", postID).Delete(&models.Like{})
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (p *PostRepositoryImpl) InsertComment(req *requests.CreateCommentReq, currentUser responses.UserResponse) (bool, error) {
	comment := &models.Comment{
		PostId: uint(req.PostID),
		UserId:      currentUser.Id,
		Content: req.Content,
	}
	result := p.db.Create(comment)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (p *PostRepositoryImpl) GetAllComment(postID int) ([]*responses.CommentResponse, error) {
	comments := []*models.Comment{}

	res := p.db.Preload("User").Order("created_at DESC").Where("post_id", postID).Find(&comments)
	var result []*responses.CommentResponse

	for _, v := range comments {

		comment := &responses.CommentResponse{
			Id:          v.Id,
			Content: v.Content,
			
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
		//	UpdatedAt: v.UpdatedAt,
		}
		result = append(result, comment)
	}

	if res.Error != nil {
		return nil, res.Error
	}
	return result, nil
}

func (p *PostRepositoryImpl) UpdateComment(req *requests.UpdateCommentReq, user responses.UserResponse) (bool, error) {
	comment := models.Comment{}
	res := p.db.First(&comment, req.CommentID)
	if res.Error != nil {

		if res.RowsAffected == 0 {
			return false, errors.New(responses.StatusResourceNotFound)
		}

		return false, res.Error
	}
	if comment.UserId != user.Id {
		return false, errors.New(responses.StatusForbidden)
	}
	comment.Content = req.Content
	result := p.db.Save(&comment)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (p *PostRepositoryImpl) DeleteComment(commentID int, user responses.UserResponse) (bool, error) {
	res := p.db.Where("user_id", user.Id).Delete(&models.Comment{}, commentID)

	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}
