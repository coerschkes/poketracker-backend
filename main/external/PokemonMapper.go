package external

import (
	"database/sql"
	"github.com/lib/pq"
	"poketracker-backend/main/domain"
)

type PokemonMapper struct {
}

func NewPokemonMapper() *PokemonMapper {
	return &PokemonMapper{}
}

func (i *PokemonMapper) MapRows(rows *sql.Rows) interface{} {
	result := make([]domain.Pokemon, 0)
	for rows.Next() {
		result = append(result, i.mapRow(rows))
	}
	return result
}

func (i *PokemonMapper) mapRow(row *sql.Rows) domain.Pokemon {
	pokemon := domain.NewPokemon()
	err := row.Scan(&pokemon.Dex, &pokemon.Name, pq.Array(&pokemon.Types), &pokemon.Shiny, &pokemon.Normal, &pokemon.Universal, &pokemon.Regional, &pokemon.UserId)
	if err != nil {
		panic(err)
	}
	return *pokemon
}
