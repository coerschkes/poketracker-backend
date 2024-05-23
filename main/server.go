package main

import (
	"github.com/labstack/echo/v4"
	"poketracker-backend/main/api"
	"poketracker-backend/main/middleware"
)

func main() {
	e := echo.New()
	loggerMiddleware := middleware.NewLoggerMiddleware()
	//add authentication when needed
	authenticationMiddleware := middleware.NewAuthenticationMiddleware()
	e.Use(loggerMiddleware.Chain, authenticationMiddleware.Chain)

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
