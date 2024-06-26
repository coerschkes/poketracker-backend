package domain

import "github.com/go-playground/validator/v10"

type (
	User struct {
		UserId    string `json:"userId" validate:"required"`
		AvatarUrl string `json:"avatarUrl" validate:"required"`
	}
	UserValidator struct {
		Validator *validator.Validate
	}
)

func (cv *UserValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}
