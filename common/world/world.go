package world

import (
	"github.com/josephnormandev/murder/common/entities/cars/dimetrodon"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/types"
)

type World struct {
	spawner *Spawner

	Dimetrodons map[types.ID]*dimetrodon.Dimetrodon
	Poles       map[types.ID]*pole.Pole // terrain elements
	Bullets     map[types.ID]*bullet.Bullet
}

func NewWorld(s *Spawner) *World {
	var game = &World{
		spawner:     s,
		Dimetrodons: map[types.ID]*dimetrodon.Dimetrodon{},
		Poles:       map[types.ID]*pole.Pole{},
		Bullets:     map[types.ID]*bullet.Bullet{},
	}
	return game
}
