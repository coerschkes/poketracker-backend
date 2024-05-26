package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"poketracker-backend/main/domain"
	"poketracker-backend/main/external"
)

type UserApi struct {
	userRepository external.UserRepository
}

func NewUserApi() *UserApi {
	return &UserApi{userRepository: external.NewUserRepositoryImpl()}
}

func (i *UserApi) RegisterRoutes(group *echo.Group) {
	group.POST("/user", i.create())
}

func (i *UserApi) create() func(c echo.Context) error {
	return func(c echo.Context) (err error) {
		u := new(domain.User)
		if err = c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(u); err != nil {
			return err
		}
		err = i.userRepository.Create(u.FirebaseUid, u.Email)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	}
}
