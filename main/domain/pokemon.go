package domain

import "github.com/go-playground/validator/v10"

type (
	Pokemon struct {
		Dex       int      `json:"dex" validate:"required"`
		Name      string   `json:"name" validate:"required"`
		Types     []string `json:"types" validate:"required"`
		Shiny     bool     `json:"shiny" validate:"required"`
		Normal    bool     `json:"normal" validate:"required"`
		Universal bool     `json:"universal" validate:"required"`
		Regional  bool     `json:"regional" validate:"required"`
		Editions  []string `json:"editions" validate:"required"`
		UserId    int
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
