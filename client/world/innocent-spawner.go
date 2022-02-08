package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/client/engine"
	"github.com/josephnormandev/murder/common/legacy-entities/innocent"
)

func (w *World) AddInnocent(i *innocent.Innocent) int {
	var id = i.GetID()

	var spawner = innocent.Spawner(w)
	i.SetSpawner(&spawner)

	var moveable = engine.Moveable(i)
	var drawable = drawer.Drawable(i)

	w.engine.AddMoveable(id, &moveable)
	w.drawer.AddDrawable(id, &drawable)
	w.Innocents[id] = i
	return id
}

func (w *World) RemoveInnocent(id int) {
	delete(w.Innocents, id)
	w.drawer.RemoveDrawable(id)
	w.engine.RemoveMoveable(id)
}

func (w *World) SpawnSword(i *innocent.Innocent) *innocent.Swingable {
	return nil
}

func (w *World) DespawnSword(id int) {

}

func (w *World) SpawnBow(i *innocent.Innocent) *innocent.Shootable {
	return nil
}

func (w *World) DespawnBow(id int) {

}
