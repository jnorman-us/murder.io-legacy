package world

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/entities/wall"
	"github.com/josephnormandev/murder/server/tcp"
)

func (w *World) AddWall(wl *wall.Wall) int {
	var id = w.NextAvailableID()
	wl.SetID(id)

	var spawn = tcp.Spawn(wl)
	var wallArrow = collisions.WallArrow(wl)
	var wallPlayer = collisions.WallPlayer(wl)

	w.network.AddSpawn(id, &spawn)
	w.collisions.AddWallArrow(id, &wallArrow)
	w.collisions.AddWallPlayer(id, &wallPlayer)
	w.Walls[id] = wl
	return id
}

func (w *World) RemoveWall(id int) {
	delete(w.Walls, id)
	w.network.RemoveSpawn(id)
	w.collisions.RemoveWallArrow(id)
	w.collisions.RemoveWallPlayer(id)
}
