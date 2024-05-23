package middleware

import "github.com/labstack/echo/v4"

type Middleware interface {
	Chain(next echo.HandlerFunc) echo.HandlerFunc
}
