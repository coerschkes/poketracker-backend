package domain

import "github.com/go-playground/validator/v10"

type (
	User struct {
		AvatarUrl string `json:"avatarUrl" validate:"required"`
		BulkMode  bool   `json:"bulkMode"`
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
