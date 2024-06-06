package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/acme/autocert"
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

	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	e.Logger.Fatal(e.StartAutoTLS(":1323"))
}

func RegisterApis(group *echo.Group, apis ...api.GenericApi) {
	for apiIndex := range apis {
		apis[apiIndex].RegisterRoutes(group)
	}
}
