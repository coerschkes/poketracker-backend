package api

import "C"
import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"poketracker-backend/main/domain"
	"poketracker-backend/main/external"
	"poketracker-backend/main/middleware"
)

type UserApi struct {
	userRepository external.UserRepository
}

func NewUserApi() *UserApi {
	return &UserApi{userRepository: external.NewUserRepositoryImpl()}
}

func (i *UserApi) RegisterRoutes(group *echo.Group) {
	group.GET("/user", i.find())
	group.POST("/user", i.create())
	group.PUT("/user", i.update())
	group.DELETE("/user", i.delete())
	group.OPTIONS("/user", i.options("GET, POST, PUT, DELETE"))
}

func (i *UserApi) options(methods string) func(c echo.Context) error {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, methods)
		c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Content-Type, Authorization")
		c.Response().WriteHeader(http.StatusOK)
		return nil
	}
}

func (i *UserApi) find() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := i.loadUserId(c)
		user, err := i.userRepository.Find(userId)

		if err != nil {
			log.Printf("user-api.find(): error while fetching user: %v\n", err)
			return err
		}
		return c.JSON(http.StatusOK, user.(domain.User))
	}
}

func (i *UserApi) create() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		p := new(domain.User)
		if err = c.Bind(p); err != nil {
			log.Printf("user-api.create(): error while binding user: %v\n", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(p); err != nil {
			log.Printf("user-api.create(): error while validating user: %v\n", err)
			return err
		}
		userId := i.loadUserId(c)
		err = i.userRepository.Create(userId, p)
		if err != nil {
			log.Printf("user-api.create(): error while creating user: %v\n", err)
			return errors.New("error while creating user")
		}
		return c.JSON(http.StatusCreated, p)
	}
}

func (i *UserApi) update() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		p := new(domain.User)
		if err = c.Bind(p); err != nil {
			log.Printf("user-api.update(): error while binding user: %v\n", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(p); err != nil {
			log.Printf("user-api.update(): error while validating user: %v\n", err)
			return err
		}
		userId := i.loadUserId(c)
		err = i.userRepository.Update(userId, p)
		if err != nil {
			log.Printf("user-api.update(): error while updating user: %v\n", err)
			return errors.New("error while updating user")
		}
		return c.JSON(http.StatusCreated, p)
	}
}

func (i *UserApi) delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := i.loadUserId(c)

		err := i.userRepository.Delete(userId)
		if err != nil {
			log.Printf("user-api.delete(): error while deleting user: %v\n", err)
			return err
		}
		return c.JSON(http.StatusOK, "user deleted ")
	}
}

func (i *UserApi) loadUserId(c echo.Context) string {
	token := c.(*middleware.AuthenticationContext).GetToken()
	return token.UID
}
