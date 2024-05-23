package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"poketracker-backend/main/middleware"
)

func main() {
	e := echo.New()
	e.Use()
	e.GET("test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	}, middleware.NewAuthenticationMiddleware().Authenticate)

	e.Logger.Fatal(e.Start(":1323"))
}
