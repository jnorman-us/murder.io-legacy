package world

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/sword"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/server/websocket"
)

func (w *World) AddSword(s *sword.Sword) int {
	var id = w.NextAvailableID()
	s.SetID(id)

	var spawn = websocket.Spawn(s)
	var tickable = logic.Tickable(s)
	var moveable = engine.Moveable(s)
	var swordPlayer = collisions.SwordPlayer(s)

	w.network.AddSpawn(id, &spawn)
	w.logic.AddTickable(id, &tickable)
	w.engine.AddMoveable(id, &moveable)
	w.collisions.AddSwordPlayer(id, &swordPlayer)
	w.Swords[id] = s
	return id
}

func (w *World) RemoveSword(id int) {
	delete(w.Swords, id)
	w.network.RemoveSpawn(id)
	w.logic.RemoveTickable(id)
	w.engine.RemoveMoveable(id)
	w.collisions.RemoveSwordPlayer(id)
}
