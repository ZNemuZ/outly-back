package usecase

import (
	"github.com/ZNemuZ/outly-back/model"
	"github.com/ZNemuZ/outly-back/repository"
	"github.com/ZNemuZ/outly-back/validator"
)

type IPostUsecase interface {
	GetAllPosts(userId uint) ([]model.PostResponce, error)
	GetPostById(userId uint, taskId uint) (model.PostResponce, error)
	CreatePost(post model.Post) (model.PostResponce, error)
	DeletePost(userId uint, postId uint) error
}
type postUsecase struct {
	pr repository.IPostRepository
	pv validator.IPostValidator
}

func NewPostUsecase(pr repository.IPostRepository, pv validator.IPostValidator) IPostUsecase {
	return &postUsecase{pr, pv}
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
			UserName:  v.User.UserName,
			UserId:    v.UserId,
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
	//バリデーション
	if err := pu.pv.PostValidate(post); err != nil {
		return model.PostResponce{}, err
	}
	userName, err := pu.pr.GetUserName(post.UserId) //userIDからuserNameを取得してくる
	if err != nil {
		return model.PostResponce{}, err
	}
	post.UserName = userName //postに取得してきたuserNameを追加
	if err := pu.pr.CreatePost(&post); err != nil {
		return model.PostResponce{}, err
	}
	resPost := model.PostResponce{
		ID:        post.ID,
		UserName:  post.UserName,
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
