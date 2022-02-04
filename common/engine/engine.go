package engine

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/collisions/collider"
)

type Engine struct {
	Moveables map[int]*Moveable
}

func NewEngine() *Engine {
	var engine = &Engine{
		Moveables: map[int]*Moveable{},
	}
	return engine
}

func (e *Engine) UpdatePhysics(time float64) {
	for _, m := range e.Moveables {
		var moveable = *m
		moveable.UpdatePosition(time)
		moveable.ClearBuffers()
	}
}

func (e *Engine) GetChannel() string {
	return "pos"
}

func (e *Engine) Flush() {

}

func (e *Engine) GetData(encoder *gob.Encoder) error {
	var colliderMap = map[int]collider.Collider{}

	for id, m := range e.Moveables {
		var moveable = *m
		colliderMap[id] = *moveable.GetCollider()
	}

	err := encoder.Encode(colliderMap)
	return err
}
