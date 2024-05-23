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
	/*
		todo:
			* check if user exists
			* retrieve all pokemon where user id is equal to the current userid
			* lazy load pokemon edition relation
			* return pokemon slice
	*/
	return nil
}

func (i *PokemonRepositoryImpl) Create(pokemon domain.Pokemon) {
	/*
		todo:
			* check if user exists
			* check if pokemon exists for user
			* create pokemon edition relation
			* return
	*/
}

func (i *PokemonRepositoryImpl) Delete(id int) {
	/*
		todo:
			* check if user exists
			* check if pokemon exists for user
			* delete pokemon edition relation
			* delete pokemon
			* return
	*/
}

func (i *PokemonRepositoryImpl) Find(id int) domain.Pokemon {
	/*
		todo:
			* check if user exists
			* check if pokemon exists for user
			* load pokemon
			* load pokemon edition relation
			* return
	*/
	return domain.Pokemon{}
}
