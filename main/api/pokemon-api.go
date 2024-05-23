package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterPokemonApiRoutes(group *echo.Group) {
	group.GET("/pokemon", func(c echo.Context) error {
		//todo: return all pokemon here
		return c.JSON(http.StatusOK, "hello world")
	})
	group.GET("/pokemon/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, "test "+id)
	})
}
