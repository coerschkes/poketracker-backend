package api

import "github.com/labstack/echo/v4"

type GenericApi interface {
	RegisterRoutes(group *echo.Group)
}
