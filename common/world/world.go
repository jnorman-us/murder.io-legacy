package world

import (
	"github.com/josephnormandev/murder/common/classes"
	"github.com/josephnormandev/murder/common/collisions"
)

// World is the struct that holds onto each subset of the objects present in this
// game instance. It is only ever mutated via the Engine so there is no need for
// channels

type World struct {
	tick              int32
	currentID         int32
	CollisionsManager *collisions.Manager
	Identifiables     map[int32]*classes.Identifiable
	Moveables         map[int32]*classes.Moveable
}

func NewWorld() *World {
	return &World{
		tick:              0,
		currentID:         0,
		CollisionsManager: collisions.NewManager(),
		Identifiables:     map[int32]*classes.Identifiable{},
		Moveables:         map[int32]*classes.Moveable{},
	}
}

func (w *World) NextAvailableID() int32 {
	w.currentID++
	return w.currentID
}

func (w *World) Tick() {
	w.tick++
}

func (w *World) GetTick() int32 {
	return w.tick
}

func (w *World) AddIdentifiable(id int32, i *classes.Identifiable) {
	w.Identifiables[id] = i
}

func (w *World) AddMoveable(id int32, m *classes.Moveable) {
	w.Moveables[id] = m
}

func (w *World) RemoveIdentifiable(id int32) {
	delete(w.Identifiables, id)
}

func (w *World) RemoveMoveable(id int32) {
	delete(w.Moveables, id)
}
