package api

import (
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"net/http"
	"poketracker-backend/main/external"
	"strconv"
)

type PokemonApi struct {
	repository external.PokemonRepository
}

func NewPokemonApi() *PokemonApi {
	return &PokemonApi{repository: external.NewPokemonRepositoryImpl()}
}

func (i *PokemonApi) RegisterRoutes(group *echo.Group) {
	group.GET("/pokemon", i.findAll())
	group.GET("/pokemon/:id", i.find())
	group.POST("/pokemon", i.create())
	group.DELETE("/pokemon/:id", i.delete())
}

func (i *PokemonApi) findAll() func(c echo.Context) error {
	return func(c echo.Context) error {
		//todo: handle error when usertoken not set
		userToken := c.Get("userToken").(*auth.Token)
		//todo: set userToken on repo
		println(userToken)
		return c.JSON(http.StatusOK, ResponseWrapper{http.StatusOK, i.repository.FindAll()})
	}
}

func (i *PokemonApi) find() func(c echo.Context) error {
	return func(c echo.Context) error {
		//todo: check id for correctness
		id := c.Param("id")
		//todo: handle error here
		parsedId, _ := strconv.Atoi(id)
		return c.JSON(http.StatusOK, ResponseWrapper{http.StatusOK, i.repository.Find(parsedId)})
	}
}

func (i *PokemonApi) create() func(c echo.Context) error {
	return func(c echo.Context) error {
		//todo read body here, create if not exists
		//todo: handle errors
		urlString := c.Request().Host + c.Request().URL.RequestURI() + "/<PokemonIdHere>"
		return c.JSON(http.StatusCreated, ResponseWrapper{http.StatusCreated, urlString})
	}
}

func (i *PokemonApi) delete() func(c echo.Context) error {
	return func(c echo.Context) error {
		//todo: check id for correctness
		id := c.Param("id")
		//todo: handle error here
		parsedId, _ := strconv.Atoi(id)
		i.repository.Delete(parsedId)
		return c.JSON(http.StatusOK, ResponseWrapper{http.StatusOK, "pokemon with id " + id + " has been deleted"})
	}
}
