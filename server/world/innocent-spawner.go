package world

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/bow"
	"github.com/josephnormandev/murder/common/entities/innocent"
	"github.com/josephnormandev/murder/common/entities/sword"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/server/websocket"
)

func (w *World) AddInnocent(i *innocent.Innocent) int {
	var id = w.NextAvailableID()
	i.SetID(id)

	var spawner = innocent.Spawner(w)
	i.SetSpawner(&spawner)

	var spawn = websocket.Spawn(i)
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
	w.Innocents[id] = i
	return id
}

func (w *World) RemoveInnocent(id int) {
	delete(w.Innocents, id)
	w.network.RemoveSpawn(id)
	w.logic.RemoveTickable(id)
	w.engine.RemoveMoveable(id)
	w.collisions.RemovePlayerArrow(id)
	w.collisions.RemovePlayerWall(id)
	w.collisions.RemovePlayerSword(id)
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
