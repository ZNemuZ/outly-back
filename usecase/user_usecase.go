package usecase

import (
	"github.com/ZNemuZ/outly-back/model"
	"github.com/ZNemuZ/outly-back/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponce, error)
	Login(user model.User) (string, error)
}
type userUsecase struct {
	ur repository.IUserRepository
}

// 依存性を注入するためのコンストラクタ
func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	//userUsecaseの構造体をurで作成してポインタを返す
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponce, error) {
	//パスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponce{}, err
	}
	//Emailとユーザーネームとハッシュ化したパスワードでuserを作成
	newUser := model.User{Email: user.Email, UserName: user.UserName, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponce{}, err
	}
	resUser := model.UserResponce{
		ID:       newUser.ID,
		Email:    newUser.Email,
		UserName: newUser.UserName,
	}
	return resUser, nil
}
