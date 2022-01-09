package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/entities/wall"
	"github.com/josephnormandev/murder/common/packet"
)

func (w *World) AddWall(wl *wall.Wall) int {
	var id int

	if w.environment.IsServer() {
		id = w.NextAvailableID()
		wl.SetID(id)

		var spawn = packet.Spawn(wl)
		var wallArrow = collisions.WallArrow(wl)
		var wallPlayer = collisions.WallPlayer(wl)

		w.network.AddSpawn(id, &spawn)
		w.collisions.AddWallArrow(id, &wallArrow)
		w.collisions.AddWallPlayer(id, &wallPlayer)
	} else {
		id = wl.GetID()
		var drawable = drawer.Drawable(wl)

		w.drawer.AddDrawable(id, &drawable)
	}
	w.Walls[id] = wl
	return id
}

func (w *World) RemoveWall(id int) {
	delete(w.Walls, id)
	if w.environment.IsServer() {
		w.network.RemoveSpawn(id)
		w.collisions.RemoveWallArrow(id)
		w.collisions.RemoveWallPlayer(id)
	} else {
		w.drawer.RemoveDrawable(id)
	}
}
