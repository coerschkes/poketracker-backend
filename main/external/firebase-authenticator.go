package external

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"strings"
)

const (
	IdentityParsingError       = 1
	ClientInstantiationError   = 2
	TokenInvalidSignatureError = 3
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
	defer i.handleError(c)
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
		if strings.Contains(fmt.Sprint(err), "failed to verify token signature") {
			panic(TokenInvalidSignatureError)
		}
		log.Fatalf("UNKNOWN TOKEN VALIDATION ERROR: %v\n", err)
	}
}

func (i *FirebaseAuthenticator) loadClient() *auth.Client {
	client, err := i.app.Auth(context.Background())
	if err != nil {
		panic(ClientInstantiationError)
	}
	return client
}

func (i *FirebaseAuthenticator) handleError(c echo.Context) {
	if r := recover(); r != nil {
		switch r {
		case IdentityParsingError:
			_ = c.JSON(http.StatusBadRequest, FirebaseError{ErrorCode: i.mapError(IdentityParsingError), Message: "Unable to parse identity object"})
		case TokenInvalidSignatureError:
			_ = c.JSON(http.StatusUnauthorized, FirebaseError{ErrorCode: i.mapError(TokenInvalidSignatureError), Message: "Token is invalid"})
		case ClientInstantiationError:
			_ = c.JSON(http.StatusInternalServerError, FirebaseError{ErrorCode: i.mapError(ClientInstantiationError), Message: "Unable to instantiate firebase client"})
			log.Fatalf("Unable to create client instance")
		}
	}
}

func (i *FirebaseAuthenticator) mapError(code int) string {
	switch code {
	case IdentityParsingError:
		return "IDENTITY_PARSING_ERROR"
	case ClientInstantiationError:
		return "CLIENT_INSTANTIATION_ERROR"
	case TokenInvalidSignatureError:
		return "TOKEN_INVALID_SIGNATURE_ERROR"
	default:
		return "UNKNOWN_ERROR"
	}
}
