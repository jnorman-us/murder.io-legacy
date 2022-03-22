package engine

import (
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
	"time"
)

type Manager struct {
	moveables       map[types.ID]*Moveable
	kinetics        map[types.ID]*engine.Kinetic
	currentlyMoving map[types.ID]bool
}

func NewManager() *Manager {
	return &Manager{
		moveables:       map[types.ID]*Moveable{},
		kinetics:        map[types.ID]*engine.Kinetic{},
		currentlyMoving: map[types.ID]bool{},
	}
}

func (m *Manager) UpdatePhysics(elapsed, total time.Duration) {
	for id, mo := range m.moveables {
		if kinetic, ok := m.kinetics[id]; ok && m.currentlyMoving[id] {
			//fmt.Println(kinetic)
			var moveable = *mo
			var timeOffset = time.Duration(0)
			var alpha = 0.0

			if elapsed >= timeOffset {
				alpha = float64(elapsed-timeOffset) / float64(total-timeOffset)
			}

			if alpha >= 0 && alpha < 1 {
				var currentPos = kinetic.StartPosition
				var futurePos = kinetic.EndPosition
				currentPos.Interpolate(futurePos, alpha)

				var currentAngle = kinetic.StartAngle
				var futureAngle = kinetic.EndAngle

				currentAngle += (futureAngle - currentAngle) * alpha
				moveable.SetPosition(currentPos)
				moveable.SetAngle(currentAngle)
			} else if alpha >= 1 {
				var futurePos = kinetic.EndPosition
				var futureAngle = kinetic.EndAngle

				moveable.SetPosition(futurePos)
				moveable.SetAngle(futureAngle)
			}
		}
		// otherwise, don't even try to move it, this object has stopped moving
	}
}

func (m *Manager) GetChannel() types.Channel {
	return 0x04
}

func (m *Manager) HandleFutureData(datums []packets.Data) {
	m.currentlyMoving = map[types.ID]bool{}
	for _, datum := range datums {
		var id = engine.GetDataID(datum)
		m.currentlyMoving[id] = true
		if kinetic, ok := m.kinetics[id]; ok {
			kinetic.FromData(datum)
		}
	}
}
