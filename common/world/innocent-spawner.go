package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/bow"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/entities/sword"
	"github.com/josephnormandev/murder/common/logic"
)

func (w *World) AddInnocent(i *innocent.Innocent) int {
	var id = w.NextAvailableID()
	var spawner = innocent.Spawner(w)
	var tickable = logic.Tickable(i)
	var drawable = drawer.Drawable(i)
	var moveable = engine.Moveable(i)
	var collidable = collisions.Collidable(i)

	i.SetID(id)
	i.SetSpawner(&spawner)

	w.Innocents[id] = i
	w.drawer.AddDrawable(id, &drawable)
	w.logic.AddTickable(id, &tickable)
	w.engine.AddMoveable(id, &moveable)
	w.collisions.AddCollidable(id, &collidable)

	return id
}

func (w *World) RemoveInnocent(id int) {
	delete(w.Innocents, id)
	w.drawer.RemoveDrawable(id)
	w.logic.RemoveTickable(id)
	w.engine.RemoveMoveable(id)
	w.collisions.RemoveCollidable(id)
}

func (w *World) SpawnSword(i *innocent.Innocent) *innocent.Swingable {
	var beholder = sword.Beholder(i)
	var s = sword.NewSword(&beholder)
	w.AddSword(s)

	var swingable = innocent.Swingable(s)

	return &swingable
}

func (w *World) DespawnSword(id int) {
	w.RemoveSword(id)
}

func (w *World) SpawnBow(i *innocent.Innocent) *innocent.Shootable {
	var holder = bow.Holder(i)
	var b = bow.NewBow(&holder)
	w.AddBow(b)

	var shootable = innocent.Shootable(b)
	return &shootable
}

func (w *World) DespawnBow(id int) {
	w.RemoveBow(id)
}
