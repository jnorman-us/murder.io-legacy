package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/bow"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/entities/sword"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/packet"
)

func (w *World) AddInnocent(i *innocent.Innocent) int {
	var id int

	var spawner = innocent.Spawner(w)
	i.SetSpawner(&spawner)

	if w.environment.IsServer() {
		id = w.NextAvailableID()
		i.SetID(id)

		var spawn = packet.Spawn(i)
		var tickable = logic.Tickable(i)
		var moveable = engine.Moveable(i)
		var playerWall = collisions.PlayerWall(i)
		var playerArrow = collisions.PlayerArrow(i)
		var playerSword = collisions.PlayerSword(i)

		w.network.AddSpawn(id, &spawn)
		w.logic.AddTickable(id, &tickable)
		w.engine.AddMoveable(id, &moveable)
		w.collisions.AddPlayerArrow(id, &playerArrow)
		w.collisions.AddPlayerWall(id, &playerWall)
		w.collisions.AddPlayerSword(id, &playerSword)
	} else {
		id = i.GetID()

		var moveable = engine.Moveable(i)
		var drawable = drawer.Drawable(i)

		w.engine.AddMoveable(id, &moveable)
		w.drawer.AddDrawable(id, &drawable)
	}
	w.Innocents[id] = i
	return id
}

func (w *World) RemoveInnocent(id int) {
	delete(w.Innocents, id)
	if w.environment.IsServer() {
		w.network.RemoveSpawn(id)
		w.logic.RemoveTickable(id)
		w.engine.RemoveMoveable(id)
		w.collisions.RemovePlayerArrow(id)
		w.collisions.RemovePlayerWall(id)
		w.collisions.RemovePlayerSword(id)
	} else {
		w.drawer.RemoveDrawable(id)
		w.engine.RemoveMoveable(id)
	}
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
