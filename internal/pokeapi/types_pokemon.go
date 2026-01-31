package pokeapi

type PokemonType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Exp  int    `json:"base_experience"`
}
