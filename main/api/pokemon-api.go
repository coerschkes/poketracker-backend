package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
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
	group.PUT("/pokemon", i.update())
	group.DELETE("/pokemon", i.deleteAll())
	group.DELETE("/pokemon/:dex", i.delete())
	group.OPTIONS("/pokemon", i.options("GET, POST, PUT, DELETE"))
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
			log.Printf("pokemon-api.findAll(): error while fetching pokemon: %v\n", err)
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
			log.Printf("pokemon-api.find(): error while fetching pokemon: %v\n", err)
			return err
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (i *PokemonApi) create() func(c echo.Context) error {
	return func(c echo.Context) (err error) {
		p := new(domain.Pokemon)
		if err = c.Bind(p); err != nil {
			log.Printf("pokemon-api.create(): error while binding pokemon: %v\n", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(p); err != nil {
			log.Printf("pokemon-api.create(): error while validating pokemon: %v\n", err)
			return err
		}
		userId := i.loadUserId(c)
		err = i.pokemonRepository.Create(*p, userId)
		if err != nil {
			log.Printf("pokemon-api.create(): error while creating pokemon: %v\n", err)
			return errors.New("error while creating pokemon")
		}
		return c.JSON(http.StatusCreated, p)
	}
}

func (i *PokemonApi) update() func(c echo.Context) error {
	return func(c echo.Context) (err error) {
		p := new(domain.Pokemon)
		if err = c.Bind(p); err != nil {
			log.Printf("pokemon-api.update(): error while binding pokemon: %v\n", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = c.Validate(p); err != nil {
			log.Printf("pokemon-api.update(): error while validating pokemon: %v\n", err)
			return err
		}
		userId := i.loadUserId(c)
		err = i.pokemonRepository.Update(*p, userId)
		if err != nil {
			log.Printf("pokemon-api.update(): error while creating pokemon: %v\n", err)
			return errors.New("error while updating pokemon")
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
			log.Printf("pokemon-api.delete(): error while deleting pokemon: %v\n", err)
			return err
		}
		return c.JSON(http.StatusOK, "pokemon with dex "+dex+" has been deleted")
	}
}

func (i *PokemonApi) loadUserId(c echo.Context) string {
	token := c.(*middleware.AuthenticationContext).GetToken()
	return token.UID
}

func (i *PokemonApi) deleteAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := i.loadUserId(c)
		err := i.pokemonRepository.DeleteAll(userId)
		if err != nil {
			log.Printf("pokemon-api.deleteAll(): error while deleting pokemon: %v\n", err)
			return err
		}
		return c.JSON(http.StatusOK, "pokemon for user "+userId+" have been deleted")
	}
}
