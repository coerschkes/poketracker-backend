package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	User struct {
		Email       string `json:"email" validate:"required,email"`
		FirebaseUid string `json:"firebaseUid" validate:"required"`
	}

	UserValidator struct {
		Validator *validator.Validate
	}
)

func (cv *UserValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
