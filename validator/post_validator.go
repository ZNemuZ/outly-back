package validator

import (
	"github.com/ZNemuZ/outly-back/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IPostValidator interface {
	PostValidate(post model.Post) error
}
type postValidator struct {
}

func NewPostValidator() IPostValidator {
	return &postValidator{}
}

func (pv *postValidator) PostValidate(post model.Post) error {
	return validation.ValidateStruct(&post,
		validation.Field(
			&post.Title,
			validation.Required.Error("title is required"),                    //必須項目
			validation.RuneLength(1, 50).Error("Please enter within 50 char"), //文字数制限
		),
		validation.Field(
			&post.Content,
			validation.Required.Error("content is required"),
			validation.RuneLength(1, 300).Error("Please enter within 300 char"),
		),
	)
}
