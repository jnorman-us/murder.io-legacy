package match

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/engine"
	"github.com/josephnormandev/murder/common/entities/cars/drifter"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
	"github.com/josephnormandev/murder/common/entities/terrain/pole"
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

func (m *Manager) SpawnPole(p *pole.Pole) types.ID {
	var id = p.ID
	p.Setup()

	var drawable = drawer.Drawable(p)

	m.drawer.AddDrawable(id, &drawable)

	return id
}

func (m *Manager) DespawnPole(id types.ID) {
	m.drawer.RemoveDrawable(id)
}

func (m *Manager) SpawnBullet(b *bullet.Bullet) types.ID {
	var id = b.ID
	b.Setup()

	var moveable = engine.Moveable(b)
	var drawable = drawer.Drawable(b)

	m.drawer.AddDrawable(id, &drawable)
	m.engine.AddMoveable(id, &moveable)

	return id
}

func (m *Manager) DespawnBullet(id types.ID) {
	m.drawer.RemoveDrawable(id)
	m.engine.RemoveMoveable(id)
}
