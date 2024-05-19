package external

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

const (
	IdentityParsingError     = 1
	ClientInstantiationError = 2
	IdTokenInvalidError      = 3
)

type FirebaseError struct {
	ErrorCode string
	Message   string
}

type ValidResponse struct {
	Message string
}

type FirebaseIdentity struct {
	IdToken string
}

type FirebaseAuthenticator struct {
	app *firebase.App
}

func NewFirebaseAuthenticator() *FirebaseAuthenticator {
	opt := option.WithCredentialsFile("config/firebase-admin-sdk-key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return &FirebaseAuthenticator{app: app}
}

func (i *FirebaseAuthenticator) Validate(c echo.Context) error {
	defer func(i *FirebaseAuthenticator, c echo.Context) {
		err := i.handleError(c)
		if err != nil {
			log.Fatalf("Unknown error: %v\n", err)
		}
	}(i, c)

	identity := i.parseFirebaseIdentity(c)
	i.verifyId(identity)

	return c.JSON(http.StatusOK, ValidResponse{Message: "Token is valid"})
}

func (i *FirebaseAuthenticator) parseFirebaseIdentity(context echo.Context) *FirebaseIdentity {
	identity := new(FirebaseIdentity)
	if err := context.Bind(identity); err != nil {
		panic(IdentityParsingError)
	}
	return identity
}

func (i *FirebaseAuthenticator) verifyId(identity *FirebaseIdentity) {
	client := i.loadClient()
	_, err := client.VerifyIDToken(context.Background(), identity.IdToken)
	if err != nil {
		panic(IdTokenInvalidError)
	}
}

func (i *FirebaseAuthenticator) loadClient() *auth.Client {
	client, err := i.app.Auth(context.Background())
	if err != nil {
		panic(ClientInstantiationError)
	}
	return client
}

func (i *FirebaseAuthenticator) handleError(c echo.Context) error {
	if r := recover(); r != nil {
		println(r)
		switch r {
		case IdentityParsingError:
			return c.JSON(http.StatusBadRequest, FirebaseError{ErrorCode: i.mapError(IdentityParsingError), Message: "Unable to parse identity object"})
		case IdTokenInvalidError:
			return c.JSON(http.StatusUnauthorized, FirebaseError{ErrorCode: i.mapError(IdTokenInvalidError), Message: "Token is invalid"})
		case ClientInstantiationError:
			log.Fatalf("Unable to create client instance")
		}
	}
	return nil
}

func (i *FirebaseAuthenticator) mapError(code int) string {
	switch code {
	case IdentityParsingError:
		return "IDENTITY_PARSING_ERROR"
	case ClientInstantiationError:
		return "CLIENT_INSTANTIATION_ERROR"
	case IdTokenInvalidError:
		return "ID_TOKEN_INVALID_ERROR"
	default:
		return "UNKNOWN_ERROR"
	}
}
