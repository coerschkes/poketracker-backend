package external

import (
	"errors"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"poketracker-backend/main/domain"
	"strconv"
)

const (
	selectPokemonQuery                 = "SELECT dex, name, types, shiny, normal, universal, regional, normalSpriteUrl, shinySpriteUrl FROM pokemon WHERE pokemon.userId = $1"
	selectPokemonByDexQuery            = "SELECT dex, name, types, shiny, normal, universal, regional FROM pokemon WHERE pokemon.userId = $1 and pokemon.dex = $2"
	selectEditionsQuery                = "SELECT editionname FROM pokemoneditionrelation WHERE pokemondexnr = $1 AND userId = $2"
	insertIntoPokemonStatement         = "INSERT INTO pokemon (dex, name, types, shiny, normal, universal, regional, userId, normalSpriteUrl, shinySpriteUrl) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	insertIntoPokemonEditionsStatement = "INSERT INTO pokemoneditionrelation (pokemondexnr, userId, editionname) VALUES ($1, $2, $3)"
	deleteFromPokemonEditionsStatement = "DELETE FROM pokemoneditionrelation WHERE pokemondexnr = $1 AND userId = $2"
	deleteFromPokemonStatement         = "DELETE FROM pokemon WHERE dex = $1 AND userId = $2"
)

type PokemonRepository interface {
	FindAll(userId string) ([]domain.Pokemon, error)
	Create(pokemon domain.Pokemon, userId string) error
	Delete(dex int, userId string) error
	Find(dex int, userId string) (domain.Pokemon, error)
}

type PokemonRepositoryImpl struct {
	connector *DatabaseConnector
}

func NewPokemonRepositoryImpl() *PokemonRepositoryImpl {
	return &PokemonRepositoryImpl{connector: NewDatabaseConnector()}
}

func (i *PokemonRepositoryImpl) FindAll(userId string) ([]domain.Pokemon, error) {
	query, err := i.connector.Query(selectPokemonQuery, NewPokemonMapper(), userId)
	if err != nil {
		return nil, errors.New("error while fetching pokemon")
	}

	pokemon := query.([]domain.Pokemon)
	for index := range pokemon {
		result, err2 := i.connector.Query(selectEditionsQuery, NewEditionMapper(), strconv.Itoa(pokemon[index].Dex), userId)
		if err2 != nil {
			return nil, errors.New("error while fetching pokemon")
		}
		pokemon[index].Editions = result.([]string)
	}
	return pokemon, nil
}

func (i *PokemonRepositoryImpl) Create(pokemon domain.Pokemon, userId string) error {
	_, err := i.Find(pokemon.Dex, userId)
	if err != nil {
		_, err := i.connector.Execute(insertIntoPokemonStatement, pokemon.Dex, pokemon.Name, pq.Array(pokemon.Types), pokemon.Shiny, pokemon.Normal, pokemon.Universal, pokemon.Regional, userId, pokemon.NormalSpriteUrl, pokemon.ShinySpriteUrl)
		if err != nil {
			return err
		}
		for index := range pokemon.Editions {
			_, err := i.connector.Execute(insertIntoPokemonEditionsStatement, pokemon.Dex, userId, pokemon.Editions[index])
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return errors.New("pokemon already exists")
	}
}

func (i *PokemonRepositoryImpl) Delete(dex int, userId string) error {
	rowsAffectedPokemon, err := i.connector.Execute(deleteFromPokemonStatement, dex, userId)
	rowsAffectedEditions, err := i.connector.Execute(deleteFromPokemonEditionsStatement, dex, userId)
	if err != nil {
		return errors.New("error while deleting pokemon")
	} else if rowsAffectedPokemon == 0 || rowsAffectedEditions == 0 {
		return errors.New("pokemon not found")
	}
	return nil
}

func (i *PokemonRepositoryImpl) Find(dex int, userId string) (domain.Pokemon, error) {
	query, err := i.connector.Query(selectPokemonByDexQuery, NewPokemonMapper(), userId, dex)
	if err != nil {
		return domain.Pokemon{}, errors.New("error while fetching pokemon")
	}
	pokemonResults := query.([]domain.Pokemon)
	if len(pokemonResults) > 0 {
		for index := range pokemonResults {
			result, err := i.connector.Query(selectEditionsQuery, NewEditionMapper(), strconv.Itoa(pokemonResults[index].Dex), userId)
			if err != nil {
				return domain.Pokemon{}, errors.New("error while fetching pokemon")
			}
			pokemonResults[index].Editions = result.([]string)
		}
		return pokemonResults[0], nil
	} else {
		return domain.Pokemon{}, errors.New("pokemon not found")
	}
}
