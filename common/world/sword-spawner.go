package world

import (
	"github.com/josephnormandev/murder/client/drawer"
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/entities/sword"
	"github.com/josephnormandev/murder/common/logic"
	"github.com/josephnormandev/murder/common/packet"
)

func (w *World) AddSword(s *sword.Sword) int {
	var id int

	if w.environment.IsServer() {
		id = w.NextAvailableID()
		s.SetID(id)

		var spawn = packet.Spawn(s)
		var tickable = logic.Tickable(s)
		var moveable = engine.Moveable(s)
		var swordPlayer = collisions.SwordPlayer(s)

		w.network.AddSpawn(id, &spawn)
		w.logic.AddTickable(id, &tickable)
		w.engine.AddMoveable(id, &moveable)
		w.collisions.AddSwordPlayer(id, &swordPlayer)
	} else {
		id = s.GetID()
		var drawable = drawer.Drawable(s)
		var moveable = engine.Moveable(s)

		w.drawer.AddDrawable(id, &drawable)
		w.engine.AddMoveable(id, &moveable)
	}
	w.Swords[id] = s
	return id
}

func (w *World) RemoveSword(id int) {
	delete(w.Swords, id)
	if w.environment.IsServer() {
		w.network.RemoveSpawn(id)
		w.logic.RemoveTickable(id)
		w.engine.RemoveMoveable(id)
		w.collisions.RemoveSwordPlayer(id)
	} else {
		w.drawer.RemoveDrawable(id)
		w.engine.RemoveMoveable(id)
	}
}
