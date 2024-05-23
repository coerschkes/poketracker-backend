package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserApi struct {
}

func NewUserApi() *UserApi {
	return &UserApi{}
}

func (i *UserApi) RegisterRoutes(group *echo.Group) {
	group.GET("/user/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.JSON(http.StatusOK, "test "+id)
	})
}
