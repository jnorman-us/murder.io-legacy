package engine

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
	"syscall/js"
	"time"
)

type Manager struct {
	kinetics  map[types.ID]packets.Kinetic
	moveables map[types.ID]*Moveable

	dataLifeStart time.Time
	dataLifeEnd   time.Time
}

func NewManager() *Manager {
	return &Manager{
		kinetics:  map[types.ID]packets.Kinetic{},
		moveables: map[types.ID]*Moveable{},
	}
}

func (m *Manager) UpdatePhysics(this js.Value, values []js.Value) interface{} {
	var currentTime = time.Now()
	var timeElapsed = currentTime.Sub(m.dataLifeStart)
	var timeTotal = m.dataLifeEnd.Sub(m.dataLifeStart)

	for id, mo := range m.moveables {
		if kinetic, ok := m.kinetics[id]; ok {
			var moveable = *mo
			var timeOffset = kinetic.GetOffsetDuration()
			var alpha = 0.0

			if timeElapsed >= timeOffset {
				alpha = float64(timeElapsed-timeOffset) / float64(timeTotal-timeOffset)
			}

			if alpha > 0 && alpha <= 1 {
				var currentPos = kinetic.GetStartPosition()
				var futurePos = kinetic.GetEndPosition()

				currentPos.Interpolate(futurePos, alpha)

				var currentAngle = kinetic.GetStartAngle()
				var futureAngle = kinetic.GetEndAngle()

				currentAngle += (futureAngle - currentAngle) * alpha

				moveable.SetPosition(currentPos)
				moveable.SetAngle(currentAngle)
			}
		}
		// otherwise, don't even try to move it, this object has stopped moving
	}
	return nil
}

func (m *Manager) GetChannel() byte {
	return 0x04
}

func (m *Manager) HandleFutureData(decoder *gob.Decoder, ttl time.Duration) error {
	var kinetics = &map[types.ID]packets.Kinetic{}

	err := decoder.Decode(kinetics)
	if err != nil {
		return err
	}

	m.dataLifeStart = time.Now()
	m.dataLifeEnd = m.dataLifeStart.Add(ttl)
	m.kinetics = *kinetics
	return nil
}
