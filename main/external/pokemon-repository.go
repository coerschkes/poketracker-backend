package external

import (
	"errors"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"poketracker-backend/main/domain"
	"strconv"
)

const (
	selectPokemonQuery                    = "SELECT dex, name, types, shiny, normal, universal, regional, normalSpriteUrl, shinySpriteUrl FROM pokemon WHERE pokemon.userId = $1"
	selectPokemonByDexQuery               = "SELECT dex, name, types, shiny, normal, universal, regional, normalspriteurl, shinyspriteurl FROM pokemon WHERE pokemon.userId = $1 and pokemon.dex = $2"
	selectEditionsQuery                   = "SELECT editionname FROM pokemoneditionrelation WHERE dex = $1 AND userId = $2"
	insertIntoPokemonStatement            = "INSERT INTO pokemon (dex, name, types, shiny, normal, universal, regional, userId, normalSpriteUrl, shinySpriteUrl) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	insertIntoPokemonEditionsStatement    = "INSERT INTO pokemoneditionrelation (dex, userId, editionname) VALUES ($1, $2, $3)"
	updatePokemonStatement                = "UPDATE pokemon SET name = $1, types = $2, shiny = $3, normal = $4, universal = $5, regional = $6, normalSpriteUrl = $7, shinySpriteUrl = $8 WHERE dex = $9 AND userId = $10"
	deleteAllFromPokemonEditionsStatement = "DELETE FROM pokemoneditionrelation WHERE userId = $1"
	deleteFromPokemonEditionsStatement    = "DELETE FROM pokemoneditionrelation WHERE dex = $1 AND userId = $2"
	deleteFromPokemonStatement            = "DELETE FROM pokemon WHERE dex = $1 AND userId = $2"
	deleteAllFromPokemonStatement         = "DELETE FROM pokemon WHERE userId = $1"
)

type PokemonRepository interface {
	FindAll(userId string) ([]domain.Pokemon, error)
	Create(pokemon domain.Pokemon, userId string) error
	Update(pokemon domain.Pokemon, userId string) error
	Delete(dex int, userId string) error
	DeleteAll(userId string) error
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
		log.Printf("pokemon-repository.FindAll(): error while fetching pokemon: %v\n", err)
		return nil, errors.New("error while fetching pokemon")
	}

	pokemon := query.([]domain.Pokemon)
	for index := range pokemon {
		result, err2 := i.connector.Query(selectEditionsQuery, NewEditionMapper(), strconv.Itoa(pokemon[index].Dex), userId)
		if err2 != nil {
			log.Printf("pokemon-repository.FindAll(): error while fetching editions: %v\n", err2)
			return nil, errors.New("error while fetching pokemon")
		}
		pokemon[index].Editions = result.([]string)
	}
	return pokemon, nil
}

func (i *PokemonRepositoryImpl) Create(pokemon domain.Pokemon, userId string) error {
	t, err := i.Find(pokemon.Dex, userId)
	log.Printf("pokemon-repository.Create(): error: %v\n", t)
	if err != nil {
		log.Printf("pokemon-repository.Create(): error while fetching pokemon: %v\n", err)
		_, err := i.connector.Execute(insertIntoPokemonStatement, pokemon.Dex, pokemon.Name, pq.Array(pokemon.Types), pokemon.Shiny, pokemon.Normal, pokemon.Universal, pokemon.Regional, userId, pokemon.NormalSpriteUrl, pokemon.ShinySpriteUrl)
		if err != nil {
			log.Printf("pokemon-repository.Create(): error while executing pokemon insert statement: %v\n", err)
			return err
		}
		for index := range pokemon.Editions {
			_, err := i.connector.Execute(insertIntoPokemonEditionsStatement, pokemon.Dex, userId, pokemon.Editions[index])
			if err != nil {
				log.Printf("pokemon-repository.Create(): error while executing editions insert statement: %v\n", err)
				return err
			}
		}
		return nil
	} else {
		return errors.New("pokemon already exists")
	}
}

func (i *PokemonRepositoryImpl) Update(pokemon domain.Pokemon, userId string) error {
	rowsAffectedPokemon, err := i.connector.Execute(updatePokemonStatement, pokemon.Name, pq.Array(pokemon.Types), pokemon.Shiny, pokemon.Normal, pokemon.Universal, pokemon.Regional, pokemon.NormalSpriteUrl, pokemon.ShinySpriteUrl, pokemon.Dex, userId)
	if err != nil {
		log.Printf("pokemon-repository.Update(): error while executing pokemon update statement: %v\n", err)
		return err
	} else if rowsAffectedPokemon == 0 {
		return errors.New("pokemon not found")
	}
	log.Printf("pokemon-repository.Update(): affected pokemon rows: %v\n", rowsAffectedPokemon)
	rowsAffectedDeletedEditions, err := i.connector.Execute(deleteFromPokemonEditionsStatement, pokemon.Dex, userId)
	if err != nil {
		log.Printf("pokemon-repository.Update(): error while deleting editions: %v\n", err)
		return err
	}
	log.Printf("pokemon-repository.Update(): affected deleted editions rows: %v\n", rowsAffectedDeletedEditions)
	for index := range pokemon.Editions {
		_, err := i.connector.Execute(insertIntoPokemonEditionsStatement, pokemon.Dex, userId, pokemon.Editions[index])
		if err != nil {
			log.Printf("pokemon-repository.Update(): error while executing editions insert statement: %v\n", err)
			return err
		}
	}
	return nil
}

func (i *PokemonRepositoryImpl) Delete(dex int, userId string) error {
	rowsAffectedPokemon, err := i.connector.Execute(deleteFromPokemonStatement, dex, userId)
	rowsAffectedEditions, err := i.connector.Execute(deleteFromPokemonEditionsStatement, dex, userId)
	if err != nil {
		log.Printf("pokemon-repository.Delete(): error while deleting pokemon: %v\n", err)
		return errors.New("error while deleting pokemon")
	} else if rowsAffectedPokemon == 0 || rowsAffectedEditions == 0 {
		return errors.New("pokemon not found")
	}
	return nil
}

func (i *PokemonRepositoryImpl) Find(dex int, userId string) (domain.Pokemon, error) {
	query, err := i.connector.Query(selectPokemonByDexQuery, NewPokemonMapper(), userId, dex)
	if err != nil {
		log.Printf("pokemon-repository.Find(): error while fetching pokemon: %v\n", err)
		return domain.Pokemon{}, errors.New("error while fetching pokemon")
	}
	pokemonResults := query.([]domain.Pokemon)
	if len(pokemonResults) > 0 {
		for index := range pokemonResults {
			result, err := i.connector.Query(selectEditionsQuery, NewEditionMapper(), strconv.Itoa(pokemonResults[index].Dex), userId)
			if err != nil {
				log.Printf("pokemon-repository.Find(): error while fetching pokemon editions: %v\n", err)
				return domain.Pokemon{}, errors.New("error while fetching pokemon")
			}
			pokemonResults[index].Editions = result.([]string)
		}
		return pokemonResults[0], nil
	} else {
		return domain.Pokemon{}, errors.New("pokemon not found")
	}
}

func (i *PokemonRepositoryImpl) DeleteAll(userId string) error {
	_, err := i.connector.Execute(deleteAllFromPokemonEditionsStatement, userId)
	if err != nil {
		log.Printf("pokemon-repository.DeleteAll(): error while deleting pokemon editions: %v\n", err)
		return errors.New("error while deleting pokemon")
	}
	_, err = i.connector.Execute(deleteAllFromPokemonStatement, userId)
	if err != nil {
		log.Printf("pokemon-repository.DeleteAll(): error while deleting pokemon: %v\n", err)
		return errors.New("error while deleting pokemon")
	}
	return nil
}
