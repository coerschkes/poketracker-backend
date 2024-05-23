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
	//authenticationMiddleware := middleware.NewAuthenticationMiddleware()
	e.Use(loggerMiddleware.Chain)

	apiGroup := e.Group("/api")
	api.RegisterPokemonApiRoutes(apiGroup)
	api.RegisterUserApiRoutes(apiGroup)

	e.Logger.Fatal(e.Start(":1323"))
}
