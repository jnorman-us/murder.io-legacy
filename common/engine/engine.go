package engine

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
)

type Engine struct {
	time *types.Time

	Moveables map[types.ID]*Moveable
	kinetics  map[types.ID]*packets.Kinetic
}

func NewEngine(ot *types.Time) *Engine {
	var engine = &Engine{
		time: ot,

		Moveables: map[types.ID]*Moveable{},
		kinetics:  map[types.ID]*packets.Kinetic{},
	}
	return engine
}

func (e *Engine) UpdatePhysics(time float64) {
	for id, m := range e.Moveables {
		var moveable = *m
		moveable.UpdatePosition(time)
		moveable.ClearBuffers()

		var col = moveable.GetCollider()
		e.kinetics[id].SetData(col.GetPosition(), col.GetAngle())
	}
}

func (e *Engine) GetChannel() byte {
	return 0x04
}

func (e *Engine) Flush() {
}

func (e *Engine) GetData(encoder *gob.Encoder) error {
	var kinetics = map[types.ID]packets.Kinetic{}

	for id, kinetic := range e.kinetics {
		if kinetic.Moved() {
			kinetics[id] = *kinetic
		}
		kinetic.Reset()
	}

	err := encoder.Encode(kinetics)
	return err
}
