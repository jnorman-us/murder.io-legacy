package engine

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
	"time"
)

type Engine struct {
	Moveables map[types.ID]*Moveable
	kinetics  map[types.ID]*packets.Kinetic

	lastSendTime time.Time
}

func NewEngine() *Engine {
	var engine = &Engine{
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
	e.lastSendTime = time.Now()

	err := encoder.Encode(kinetics)
	return err
}
