package engine

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/types"
)

type Engine struct {
	Moveables map[types.ID]*Moveable
	kinetics  map[types.ID]*Kinetic
}

func NewEngine() *Engine {
	var engine = &Engine{
		Moveables: map[types.ID]*Moveable{},
		kinetics:  map[types.ID]*Kinetic{},
	}
	return engine
}

func (e *Engine) UpdatePhysics(time float64) {
	for id, m := range e.Moveables {
		var moveable = *m
		moveable.UpdatePosition(time)
		moveable.ClearBuffers()

		var position = moveable.GetPosition()
		var angle = moveable.GetAngle()
		e.kinetics[id].Set(position, angle)
	}
}

func (e *Engine) GetChannel() types.Channel {
	return 0x04
}

func (e *Engine) GetData() []data.Data {
	var datums []data.Data

	for _, k := range e.kinetics {
		if k.Moved() {
			datums = append(datums, k.GetData())
		}
		k.Restart()
	}
	return datums
}
