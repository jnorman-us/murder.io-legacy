package world

import (
	"github.com/josephnormandev/murder/common/entities/cars/dimetrodon"
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
	"github.com/josephnormandev/murder/common/types"
)

type Spawner interface {
	SpawnDrifter(*drifter.Drifter) types.ID
	SpawnPole(*pole.Pole) types.ID
	SpawnBullet(*bullet.Bullet) types.ID
	SpawnDimetrodon(*dimetrodon.Dimetrodon) types.ID

	DespawnDrifter(types.ID)
	DespawnDimetrodon(types.ID)
	DespawnPole(types.ID)
	DespawnBullet(types.ID)
}
