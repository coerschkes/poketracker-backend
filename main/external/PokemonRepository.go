package external

import "poketracker-backend/main/domain"

type PokemonRepository interface {
	FindAll() []domain.Pokemon
	Create(pokemon domain.Pokemon)
	Delete(id int)
	Find(id int) domain.Pokemon
}

type PokemonRepositoryImpl struct {
	UserId int
}

func NewPokemonRepositoryImpl() *PokemonRepositoryImpl {
	return &PokemonRepositoryImpl{}
}

func (i *PokemonRepositoryImpl) FindAll() []domain.Pokemon {
	//todo: implement
	return nil
}

func (i *PokemonRepositoryImpl) Create(pokemon domain.Pokemon) {
	//todo: implement
}

func (i *PokemonRepositoryImpl) Delete(id int) {
	//todo: implement
}

func (i *PokemonRepositoryImpl) Find(id int) domain.Pokemon {
	//todo: implement
	return domain.Pokemon{}
}
