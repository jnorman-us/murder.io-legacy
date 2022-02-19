package world

import (
	"github.com/josephnormandev/murder/common/types"
)

func (w *World) DisableBullet(id types.ID) {
	(*w.spawner).DisableBullet(id)
}
