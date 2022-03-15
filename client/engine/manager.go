package engine

import (
	"github.com/josephnormandev/murder/common/types"
	"time"
)

type Manager struct {
	moveables map[types.ID]*Moveable
}

func NewManager() *Manager {
	return &Manager{
		moveables: map[types.ID]*Moveable{},
	}
}

func (m *Manager) UpdatePhysics(elapsed, total time.Duration) { /*
		for id, mo := range m.moveables {
			if kinetic, ok := m.kinetics[id]; ok {
				//fmt.Println(kinetic)
				var moveable = *mo
				var timeOffset = kinetic.Offset
				var alpha = 0.0

				if elapsed >= timeOffset {
					alpha = float64(elapsed-timeOffset) / float64(total-timeOffset)
				}

				if alpha > 0 && alpha < 1 {
					var currentPos = kinetic.GetStartPosition()
					var futurePos = kinetic.GetEndPosition()

					currentPos.Interpolate(futurePos, alpha)

					var currentAngle = kinetic.GetStartAngle()
					var futureAngle = kinetic.GetEndAngle()

					currentAngle += (futureAngle - currentAngle) * alpha
					moveable.SetPosition(currentPos)
					moveable.SetAngle(currentAngle)
				} else if alpha >= 1 {
					var futurePos = kinetic.GetEndPosition()
					var futureAngle = kinetic.GetEndAngle()

					moveable.SetPosition(futurePos)
					moveable.SetAngle(futureAngle)
				}
			}
			// otherwise, don't even try to move it, this object has stopped moving
		}*/
}
