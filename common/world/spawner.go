package world

import (
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/types"
)

type Spawner interface {
	SpawnDrifter(d *drifter.Drifter) types.ID

	DespawnDrifter(types.ID)
}
