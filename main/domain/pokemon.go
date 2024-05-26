package domain

type Pokemon struct {
	Dex       int
	Name      string
	Types     []string
	Shiny     bool
	Normal    bool
	Universal bool
	Regional  bool
	Editions  []string
	UserId    int
}

func NewPokemon() *Pokemon {
	return &Pokemon{
		Dex:       0,
		Name:      "",
		Types:     []string{},
		Shiny:     false,
		Normal:    false,
		Universal: false,
		Regional:  false,
		//todo: care: might cause nil pointer exception
		Editions: nil,
		UserId:   0,
	}
}
