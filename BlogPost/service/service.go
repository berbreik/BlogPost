// service/service.go

package service

import (
	"BlogPost/model"
	"errors"
	"github.com/jinzhu/gorm"
)

// BlogPostService is responsible for managing blog posts.
type BlogPostService interface {
	GetPostByID(id uint) (*model.BlogPost, error)
	GetAllPosts() ([]model.BlogPost, error)
	CreatePost(post *model.BlogPost) error
	UpdatePost(id uint, updatedPost *model.BlogPost) error
	DeletePost(id uint) error
}

// BlogPostServiceImpl implements BlogPostService.
type BlogPostServiceImpl struct {
	DB *gorm.DB
}

// NewBlogPostService creates a new BlogPostService instance.
func NewBlogPostService(db *gorm.DB) BlogPostService {
	return &BlogPostServiceImpl{DB: db}
}

var _ BlogPostService = &BlogPostServiceImpl{} // Ensure implementation

var ErrPostNotFound = errors.New("post not found")

func (s *BlogPostServiceImpl) GetPostByID(id uint) (*model.BlogPost, error) {
	post := &model.BlogPost{}
	if err := s.DB.First(post, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrPostNotFound
		}
		return nil, err
	}
	return post, nil
}

func (s *BlogPostServiceImpl) GetAllPosts() ([]model.BlogPost, error) {
	posts := []model.BlogPost{}
	if err := s.DB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *BlogPostServiceImpl) CreatePost(post *model.BlogPost) error {
	if err := s.DB.Create(post).Error; err != nil {
		return err
	}
	return nil
}

func (s *BlogPostServiceImpl) UpdatePost(id uint, updatedPost *model.BlogPost) error {
	if err := s.DB.First(updatedPost, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrPostNotFound
		}
		return err
	}
	if err := s.DB.Save(updatedPost).Error; err != nil {
		return err
	}
	return nil
}

func (s *BlogPostServiceImpl) DeletePost(id uint) error {
	if err := s.DB.Delete(&model.BlogPost{}, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return ErrPostNotFound
		}
		return err
	}
	return nil
}
