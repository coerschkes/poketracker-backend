package middleware

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"log"
	"strings"
)

type ErrorMessage struct {
	Message string
}

type AuthenticationMiddleware struct {
	firebaseAdminSdkClient *auth.Client
}

func NewAuthenticationMiddleware() *AuthenticationMiddleware {
	opt := option.WithCredentialsFile("../config/firebase-admin-sdk-key.json")
	app, errApp := firebase.NewApp(context.Background(), nil, opt)
	if errApp != nil {
		log.Fatalf("error initializing app: %v\n", errApp)
	}
	client, errClient := app.Auth(context.Background())
	if errClient != nil {
		log.Fatalf("error instantiating firebaseAdminSdkClient: %v\n", errClient)
	}
	return &AuthenticationMiddleware{firebaseAdminSdkClient: client}
}

func (i *AuthenticationMiddleware) Chain(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := i.parseAuthorizationHeader(c)
		if token == "" {
			return c.JSON(400, ErrorMessage{"Unable to parse authorization header}"})
		}
		responseToken, err := i.verifyToken(token)
		if err != nil {
			return c.JSON(401, ErrorMessage{"Token is invalid"})
		}
		c.Set("userToken", responseToken)
		return next(c)
	}
}

func (i *AuthenticationMiddleware) parseAuthorizationHeader(c echo.Context) string {
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	splitHeader := strings.Split(authHeader, " ")
	if len(splitHeader) != 2 || splitHeader[0] != "Bearer" {
		log.Printf("Unable to parse auth header: '%v'\n", authHeader)
		return ""
	} else {
		return splitHeader[1]
	}
}

func (i *AuthenticationMiddleware) verifyToken(token string) (*auth.Token, error) {
	responseToken, err := i.firebaseAdminSdkClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		if strings.Contains(fmt.Sprint(err), "failed to verify token signature") {
			log.Printf("unable to verify token: '%v'\n", err)
		} else {
			log.Printf("unknown token validation error: '%v'\n", err)
		}
		return nil, err
	}
	return responseToken, nil
}
