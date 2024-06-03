package domain

import (
	"github.com/go-playground/validator/v10"
)

type (
	Pokemon struct {
		Dex             int      `json:"dex" validate:"required"`
		Name            string   `json:"name" validate:"required"`
		Types           []string `json:"types" validate:"required"`
		Shiny           bool     `json:"shiny"`
		Normal          bool     `json:"normal"`
		Universal       bool     `json:"universal"`
		Regional        bool     `json:"regional"`
		Editions        []string `json:"editions" validate:"required"`
		NormalSpriteUrl string   `json:"normalSpriteUrl" validate:"required"`
		ShinySpriteUrl  string   `json:"shinySpriteUrl" validate:"required"`
	}
	PokemonValidator struct {
		Validator *validator.Validate
	}
)

func (cv *PokemonValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}
