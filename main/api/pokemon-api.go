package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"poketracker-backend/main/domain"
	"poketracker-backend/main/external"
	"poketracker-backend/main/middleware"
	"strconv"
)

type PokemonApi struct {
	pokemonRepository external.PokemonRepository
}

func NewPokemonApi() *PokemonApi {
	return &PokemonApi{pokemonRepository: external.NewPokemonRepositoryImpl()}
}

func (i *PokemonApi) RegisterRoutes(group *echo.Group) {
	group.GET("/pokemon", i.findAll())
	group.GET("/pokemon/:dex", i.find())
	group.POST("/pokemon", i.create())
	group.DELETE("/pokemon/:dex", i.delete())
	group.OPTIONS("/pokemon", i.options("GET, POST"))
	group.OPTIONS("/pokemon/:dex", i.options("GET, DELETE"))
}

func (i *PokemonApi) options(methods string) func(c echo.Context) error {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, methods)
		c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Content-Type, Authorization")
		c.Response().WriteHeader(http.StatusOK)
		return nil
	}
}

func (i *PokemonApi) findAll() func(c echo.Context) error {
	return func(c echo.Context) error {
		userId := i.loadUserId(c)

		result, err := i.pokemonRepository.FindAll(userId)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (i *PokemonApi) find() func(c echo.Context) error {
	return func(c echo.Context) error {
		userId := i.loadUserId(c)

		dex := c.Param("dex")
		parsedId, _ := strconv.Atoi(dex)
		result, err := i.pokemonRepository.Find(parsedId, userId)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (i *PokemonApi) create() func(c echo.Context) error {
	return func(c echo.Context) (err error) {
		p := new(domain.Pokemon)
		if err = c.Bind(p); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(p); err != nil {
			return err
		}
		userId := i.loadUserId(c)
		err = i.pokemonRepository.Create(*p, userId)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, p)
	}
}

func (i *PokemonApi) delete() func(c echo.Context) error {
	return func(c echo.Context) error {
		dex := c.Param("dex")
		userId := i.loadUserId(c)

		parsedDex, _ := strconv.Atoi(dex)
		err := i.pokemonRepository.Delete(parsedDex, userId)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, "pokemon with dex "+dex+" has been deleted")
	}
}

func (i *PokemonApi) loadUserId(c echo.Context) string {
	token := c.(*middleware.AuthenticationContext).GetToken()
	return token.UID
}
