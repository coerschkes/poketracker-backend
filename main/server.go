package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"poketracker-backend/main/api"
	"poketracker-backend/main/domain"
	"poketracker-backend/main/middleware"
)

func main() {
	e := echo.New()
	e.HTTPErrorHandler = middleware.NewHttpErrorHandler().HandleError
	e.Validator = &domain.PokemonValidator{Validator: validator.New()}
	loggerMiddleware := middleware.NewLoggerMiddleware()
	authenticationMiddleware := middleware.NewAuthenticationMiddleware()
	headerMiddleware := middleware.NewHeaderMiddleware()
	e.Use(loggerMiddleware.Chain, headerMiddleware.Chain, authenticationMiddleware.Chain)

	apiGroup := e.Group("/api")
	pokemonApi := api.NewPokemonApi()
	RegisterApis(apiGroup, pokemonApi)

	if err := e.StartTLS(":1323", "../config/server.crt", "../config/server.key"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func RegisterApis(group *echo.Group, apis ...api.GenericApi) {
	for apiIndex := range apis {
		apis[apiIndex].RegisterRoutes(group)
	}
}
