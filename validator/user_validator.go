package validator

import (
	"github.com/ZNemuZ/outly-back/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidator(user model.User) error
}

type userValidator struct {
}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidator(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("Please enter within 30 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.UserName,
			validation.Required.Error("userName is required"),
			validation.RuneLength(5, 30).Error("Please enter between 5 and 30 char"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("Please enter between 6 and 30 char"),
		),
	)
}
