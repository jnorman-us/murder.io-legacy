package entities

import (
	"github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/types"
)

type Player struct {
	Entity
}

func NewPlayer() *Player {
	var player = &Player{}

	player.Setup(
		[]collider.Rectangle{},
		[]collider.Circle{
			collider.NewCircle(types.NewVector(0, 0), 2),
		},
	)
	player.SetStats(20, 5)

	return player
}

func (p *Player) HitByZombie(z *collisions.ZombiePlayerCollidable) {

}