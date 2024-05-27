package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"poketracker-backend/main/api"
	"poketracker-backend/main/domain"
	"poketracker-backend/main/middleware"
)

func main() {
	e := echo.New()
	e.HTTPErrorHandler = middleware.NewHttpErrorHandler().HandleError
	e.Validator = &domain.UserValidator{Validator: validator.New()}
	e.Validator = &domain.PokemonValidator{Validator: validator.New()}
	loggerMiddleware := middleware.NewLoggerMiddleware()
	authenticationMiddleware := middleware.NewAuthenticationMiddleware()
	headerMiddleware := middleware.NewHeaderMiddleware()
	e.Use(loggerMiddleware.Chain, authenticationMiddleware.Chain, headerMiddleware.Chain)

	apiGroup := e.Group("/api")
	pokemonApi := api.NewPokemonApi()
	userApi := api.NewUserApi()
	RegisterApis(apiGroup, pokemonApi, userApi)

	e.Logger.Fatal(e.Start(":1323"))
}

func RegisterApis(group *echo.Group, apis ...api.GenericApi) {
	for apiIndex := range apis {
		apis[apiIndex].RegisterRoutes(group)
	}
}
