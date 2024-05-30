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

func (i *PokemonMapper) MapRows(rows *sql.Rows) (interface{}, error) {
	result := make([]domain.Pokemon, 0)
	for rows.Next() {
		row, err := i.mapRow(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func (i *PokemonMapper) mapRow(row *sql.Rows) (domain.Pokemon, error) {
	pokemon := new(domain.Pokemon)
	err := row.Scan(&pokemon.Dex, &pokemon.Name, pq.Array(&pokemon.Types), &pokemon.Shiny, &pokemon.Normal, &pokemon.Universal, &pokemon.Regional)
	if err != nil {
		return domain.Pokemon{}, err
	}
	return *pokemon, err
}
