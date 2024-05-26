package external

import (
	_ "github.com/lib/pq"
	"poketracker-backend/main/domain"
	"strconv"
)

//todo: https://echo.labstack.com/docs/context for user information retrieval (custom context?)

const (
	selectPokemonQuery  = "SELECT * FROM pokemon WHERE pokemon.userId = $1"
	selectEditionsQuery = "SELECT editionname FROM pokemoneditionrelation WHERE pokemondexnr = $1 AND userId = $2"
)

type PokemonRepository interface {
	FindAll() ([]domain.Pokemon, error)
	Create(pokemon domain.Pokemon) error
	Delete(id int) error
	Find(id int) (domain.Pokemon, error)
}

type PokemonRepositoryImpl struct {
	UserId    int
	Connector *DatabaseConnector
}

func NewPokemonRepositoryImpl(userId int) *PokemonRepositoryImpl {
	return &PokemonRepositoryImpl{UserId: userId, Connector: NewDatabaseConnector()}
}

func (i *PokemonRepositoryImpl) FindAll() ([]domain.Pokemon, error) {
	var err error = nil
	defer i.recoverError(&err)

	pokemonResults := i.Connector.Query(selectPokemonQuery, NewPokemonMapper(), strconv.Itoa(i.UserId)).([]domain.Pokemon)
	for index := range pokemonResults {
		result := i.Connector.Query(selectEditionsQuery, NewEditionMapper(), strconv.Itoa(pokemonResults[index].Dex), strconv.Itoa(i.UserId))
		pokemonResults[index].Editions = result.([]string)
	}
	return pokemonResults, err
}

func (i *PokemonRepositoryImpl) Create(pokemon domain.Pokemon) error {
	return nil
	/*
		todo:
			* check if user exists
			* check if pokemon exists for user
			* create pokemon edition relation
			* return
	*/
}

func (i *PokemonRepositoryImpl) Delete(id int) error {
	return nil
	/*
		todo:
			* check if user exists
			* check if pokemon exists for user
			* delete pokemon edition relation
			* delete pokemon
			* return
	*/
}

func (i *PokemonRepositoryImpl) Find(id int) (domain.Pokemon, error) {
	/*
		todo:
			* check if user exists
			* check if pokemon exists for user
			* load pokemon
			* load pokemon edition relation
			* return
	*/
	return domain.Pokemon{}, nil
}

func (*PokemonRepositoryImpl) recoverError(err *error) {
	if r := recover(); r != nil {
		*err = r.(error)
	}
}
