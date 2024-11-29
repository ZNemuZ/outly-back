package repository

import "github.com/ZNemuZ/outly-back/model"

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}
