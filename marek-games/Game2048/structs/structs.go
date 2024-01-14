package structs

type Tile struct {
	Value  int
	Merged bool
}

type GameState struct {
	Grid     [][]*Tile
	Score    int
	GameOver bool
}
