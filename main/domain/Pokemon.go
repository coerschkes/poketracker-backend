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
}
