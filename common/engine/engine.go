package engine

import (
	"encoding/gob"
	collider2 "github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/types"
)

type Engine struct {
	Moveables map[types.ID]*Moveable
}

func NewEngine() *Engine {
	var engine = &Engine{
		Moveables: map[types.ID]*Moveable{},
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

func (e *Engine) GetChannel() byte {
	return 0x04
}

func (e *Engine) Flush() {

}

func (e *Engine) GetData(encoder *gob.Encoder) error {
	var colliderMap = map[types.ID]collider2.Collider{}

	for id, m := range e.Moveables {
		var moveable = *m
		//if moveable.Changed() {
		colliderMap[id] = *moveable.GetCollider()
		//	moveable.Reset()
		//}
	}

	err := encoder.Encode(colliderMap)
	return err
}
