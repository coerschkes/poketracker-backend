package middleware

import (
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
)

type AuthenticationContext struct {
	echo.Context
	*auth.Token
}

func (c *AuthenticationContext) SetToken(token *auth.Token) {
	c.Token = token
}

func (c *AuthenticationContext) GetToken() *auth.Token {
	return c.Token
}
