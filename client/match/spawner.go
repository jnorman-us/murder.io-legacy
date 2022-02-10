package match

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/engine"
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/types"
)

func (m *Manager) SpawnDrifter(d *drifter.Drifter) types.ID {
	var id = d.ID
	d.Setup()

	if d.GetUserID() == m.Username {
		var centerable = drawer.Centerable(d)
		m.drawer.SetCenterable(&centerable)
	}

	var drawable = drawer.Drawable(d)
	var moveable = engine.Moveable(d)

	m.drawer.AddDrawable(id, &drawable)
	m.engine.AddMoveable(id, &moveable)

	return id
}

func (m *Manager) DespawnDrifter(id types.ID) {
	m.drawer.RemoveDrawable(id)
	m.engine.RemoveMoveable(id)
}
