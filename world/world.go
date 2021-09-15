package world

import (
	"github.com/josephnormandev/murder/classes"
	"github.com/josephnormandev/murder/collisions"
)

// World is the struct that holds onto each subset of the objects present in this
// game instance. It is only ever mutated via the Engine so there is no need for
// channels
type World struct {
	currentID         int32
	CollisionsManager *collisions.Manager
	Identifiables     map[int32]*classes.Identifiable
	Moveables         map[int32]*classes.Moveable
	Drawables         map[int32]*classes.Drawable
}

func NewWorld() *World {
	return &World{
		currentID:         0,
		CollisionsManager: collisions.NewManager(),
		Identifiables:     map[int32]*classes.Identifiable{},
		Moveables:         map[int32]*classes.Moveable{},
		Drawables:         map[int32]*classes.Drawable{},
	}
}

func (w *World) NextAvailableID() int32 {
	w.currentID++
	return w.currentID
}

func (w *World) AddIdentifiable(id int32, i *classes.Identifiable) {
	w.Identifiables[id] = i
}

func (w *World) AddDrawable(id int32, d *classes.Drawable) {
	w.Drawables[id] = d
}

func (w *World) AddMoveable(id int32, m *classes.Moveable) {
	w.Moveables[id] = m
}

func (w *World) RemoveIdentifiable(id int32) {
	delete(w.Identifiables, id)
}

func (w *World) RemoveDrawable(id int32) {
	delete(w.Drawables, id)
}

func (w *World) RemoveMoveable(id int32) {
	delete(w.Moveables, id)
}
