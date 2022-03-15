package engine

import (
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
