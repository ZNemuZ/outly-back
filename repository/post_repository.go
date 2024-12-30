package repository

import (
	"fmt"

	"github.com/ZNemuZ/outly-back/model"
	"gorm.io/gorm"
)

type IPostRepository interface {
	GetAllPosts(post *[]model.Post, userId uint) error
	GetPostById(post *model.Post, userId uint, postId uint) error
	CreatePost(post *model.Post) error
	DeletePost(userId uint, postId uint) error
	GetUserName(userId uint) (string, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) IPostRepository {
	return &postRepository{db}
}

func (pr *postRepository) GetAllPosts(posts *[]model.Post, userId uint) error {
	if err := pr.db.Joins("User").Order("created_at DESC").Find(posts).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) GetPostById(post *model.Post, userId uint, postId uint) error {
	if err := pr.db.Joins("User").Where("user_id=?", userId).First(post, postId).Error; err != nil {
		return err
	}
	return nil
}
func (pr *postRepository) CreatePost(post *model.Post) error {
	if err := pr.db.Create(post).Error; err != nil {
		return err
	}
	return nil
}

func (pr *postRepository) DeletePost(userId uint, postId uint) error {
	result := pr.db.Where("id=? AND user_id=?", postId, userId).Delete(&model.Post{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
func (pr *postRepository) GetUserName(userId uint) (string, error) {
	var user model.User
	if err := pr.db.First(&user, userId).Error; err != nil {
		return "", err
	}
	return user.UserName, nil
}
