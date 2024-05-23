package external

import (
	_ "github.com/lib/pq"
	"poketracker-backend/main/domain"
	"strconv"
)

//todo: install go get github.com/lib/pq for postgres communication https://pkg.go.dev/github.com/lib/pq#section-readme
//todo: https://echo.labstack.com/docs/context for user information retrieval (custom context?)

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
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	mapper := NewPokemonMapper()
	//todo: load pokemon edition relation
	result := i.Connector.Query("SELECT * FROM pokemon WHERE pokemon.userId = $1", mapper, strconv.Itoa(i.UserId))
	return result.([]domain.Pokemon), err
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
