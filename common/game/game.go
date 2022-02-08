package game

import (
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/types"
)

type Game struct {
	Players map[types.UserID]int

	deletions *Deletions
	Drifters  map[types.ID]*drifter.Drifter // cars
	Poles     map[types.ID]*pole.Pole       // terrain elements
}

func NewGame() *Game {
	var game = &Game{
		Players: map[types.UserID]int{},

		deletions: NewDeletions(),
		Drifters:  map[types.ID]*drifter.Drifter{},
		Poles:     map[types.ID]*pole.Pole{},
	}
	return game
}

func (g *Game) Deletions() *Deletions {
	return g.deletions
}
