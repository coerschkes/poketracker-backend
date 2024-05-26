package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"poketracker-backend/main/external"
	"poketracker-backend/main/middleware"
	"strconv"
)

type PokemonApi struct {
	pokemonRepository external.PokemonRepository
	userRepository    external.UserRepository
}

func NewPokemonApi() *PokemonApi {
	return &PokemonApi{pokemonRepository: external.NewPokemonRepositoryImpl(), userRepository: external.NewUserRepositoryImpl()}
}

func (i *PokemonApi) RegisterRoutes(group *echo.Group) {
	group.GET("/pokemon", i.findAll())
	group.GET("/pokemon/:dex", i.find())
	group.POST("/pokemon", i.create())
	group.DELETE("/pokemon/:dex", i.delete())
}

func (i *PokemonApi) findAll() func(c echo.Context) error {
	return func(c echo.Context) error {
		userId, err := i.loadUserId(c)
		if err != nil {
			return err
		}

		result, err := i.pokemonRepository.FindAll(userId)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, ResponseWrapper{http.StatusOK, result})
	}
}

func (i *PokemonApi) find() func(c echo.Context) error {
	return func(c echo.Context) error {
		userId, err := i.loadUserId(c)
		if err != nil {
			return err
		}

		dex := c.Param("dex")
		parsedId, _ := strconv.Atoi(dex)
		result, err := i.pokemonRepository.Find(parsedId, userId)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, ResponseWrapper{http.StatusOK, result})
	}
}

func (i *PokemonApi) create() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(http.StatusBadRequest, ResponseWrapper{http.StatusNotFound, "not implemented yet "})
	}
}

func (i *PokemonApi) delete() func(c echo.Context) error {
	return func(c echo.Context) error {
		dex := c.Param("dex")
		userId, err := i.loadUserId(c)
		if err != nil {
			return err
		}

		parsedDex, _ := strconv.Atoi(dex)
		err = i.pokemonRepository.Delete(parsedDex, userId)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, ResponseWrapper{http.StatusOK, "pokemon with dex " + dex + " has been deleted"})
	}
}

func (i *PokemonApi) loadUserId(c echo.Context) (int, error) {
	token := c.(*middleware.AuthenticationContext).GetToken()
	result, err := i.userRepository.Find(token.UID)
	return result, err
}
