package world

import (
	"github.com/josephnormandev/murder/common/entities/cars/dimetrodon"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/types"
)

type Spawner interface {
	SpawnPole(*pole.Pole) types.ID
	SpawnBullet(*bullet.Bullet) types.ID
	SpawnDimetrodon(*dimetrodon.Dimetrodon) types.ID

	DespawnDimetrodon(types.ID)
	DespawnPole(types.ID)
	DespawnBullet(types.ID)

	DisableBullet(types.ID)
}
