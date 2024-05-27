package middleware

import (
	"github.com/labstack/echo/v4"
)

type HeaderMiddleware struct {
}

func NewHeaderMiddleware() *HeaderMiddleware {
	return &HeaderMiddleware{}
}

func (i *HeaderMiddleware) Chain(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
		return next(c)
	}
}
