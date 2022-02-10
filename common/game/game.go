package game

import "github.com/josephnormandev/murder/common/types"

type Game struct {
	Players map[types.UserID]int
	// other Game State contents
}

func NewGame() *Game {
	return &Game{
		Players: map[types.UserID]int{},
	}
}
