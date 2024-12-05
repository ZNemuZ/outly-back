package usecase

import (
	"os"
	"time"

	"github.com/ZNemuZ/outly-back/model"
	"github.com/ZNemuZ/outly-back/repository"
	"github.com/ZNemuZ/outly-back/validator"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponce, error)
	Login(user model.User) (string, error)
}
type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

// 依存性を注入するためのコンストラクタ
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	//userUsecaseの構造体をurで作成してポインタを返す
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponce, error) {
	//バリデーション
	if err := uu.uv.SignUpValidator(user); err != nil {
		return model.UserResponce{}, err
	}
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

func (uu *userUsecase) Login(user model.User) (string, error) {
	//バリデーション
	if err := uu.uv.LoginValidator(user); err != nil {
		return "", err
	}
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	//パスワードの検証
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	//jwtトークンの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(), //jwtの有効期限
	})
	//署名
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
