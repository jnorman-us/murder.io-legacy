package game

import (
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/types"
)

type Game struct {
	entityID types.ID

	players map[types.UserID]int

	// cars
	drifters map[types.ID]*drifter.Drifter

	// terrain elements
	poles map[types.ID]*pole.Pole
}

func NewGame() *Game {
	var game = &Game{
		drifters: map[types.ID]*drifter.Drifter{},

		poles: map[types.ID]*pole.Pole{},
	}
	return game
}
