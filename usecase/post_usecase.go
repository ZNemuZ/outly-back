package usecase

import (
	"github.com/ZNemuZ/outly-back/model"
	"github.com/ZNemuZ/outly-back/repository"
)

type IPostUsecase interface {
	GetAllPosts(userId uint) ([]model.PostResponce, error)
	GetPostById(userId uint, taskId uint) (model.PostResponce, error)
	CreatePost(post model.Post) (model.PostResponce, error)
	DeletePost(userId uint, postId uint) error
}
type postUsecase struct {
	pr repository.IPostRepository
}

func NewPostUsecase(pr repository.IPostRepository) IPostUsecase {
	return &postUsecase{pr}
}

func (pu *postUsecase) GetAllPosts(userId uint) ([]model.PostResponce, error) {
	posts := []model.Post{}
	if err := pu.pr.GetAllPosts(&posts, userId); err != nil {
		return nil, err
	}
	resPosts := []model.PostResponce{}
	for _, v := range posts {
		p := model.PostResponce{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			NiceCount: v.NiceCount,
			CreatedAt: v.CreatedAt,
		}
		resPosts = append(resPosts, p)
	}
	return resPosts, nil
}
func (pu *postUsecase) GetPostById(userId uint, postId uint) (model.PostResponce, error) {
	post := model.Post{}
	if err := pu.pr.GetPostById(&post, userId, postId); err != nil {
		return model.PostResponce{}, err
	}
	resPost := model.PostResponce{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		NiceCount: post.NiceCount,
		CreatedAt: post.CreatedAt,
	}
	return resPost, nil
}

func (pu *postUsecase) CreatePost(post model.Post) (model.PostResponce, error) {
	if err := pu.pr.CreatePost(&post); err != nil {
		return model.PostResponce{}, err
	}
	resPost := model.PostResponce{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		NiceCount: post.NiceCount,
		CreatedAt: post.CreatedAt,
	}
	return resPost, nil
}

func (pu postUsecase) DeletePost(userId uint, postId uint) error {
	if err := pu.pr.DeletePost(userId, postId); err != nil {
		return err
	}
	return nil
}
