package engine

import (
	"encoding/gob"
	"fmt"
	"github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/types"
	"math"
	"time"
)

type Manager struct {
	colliders map[types.ID]collider.Collider
	moveables map[types.ID]*Moveable

	dataLifeStart time.Time
	dataLifeEnd   time.Time
}

func NewManager() *Manager {
	return &Manager{
		colliders: map[types.ID]collider.Collider{},
		moveables: map[types.ID]*Moveable{},
	}
}

func (m *Manager) UpdatePhysics(ms time.Duration) {
	var currentTime = time.Now()
	var timeElapsed = currentTime.Sub(m.dataLifeStart)
	var timeTotal = m.dataLifeEnd.Sub(m.dataLifeStart)

	var alpha = 0.0
	if timeTotal != 0 {
		alpha = float64(timeElapsed) / float64(timeTotal)
		alpha = math.Min(1, alpha)
	}

	if alpha >= 1 {
		/*for _, mo := range m.moveables {
			// we seem to be missing packets, let's extrapolate...
			var moveable = *mo
			moveable.UpdatePosition(ms / (1000 / 60))
		}*/
	} else {
		for id, mo := range m.moveables {
			if collider, ok := m.colliders[id]; ok {
				var moveable = *mo
				var currentPos = moveable.GetPosition()
				var futurePos = collider.GetPosition()

				currentPos.Interpolate(futurePos, alpha)

				var currentAngle = moveable.GetAngle()
				var futureAngle = collider.GetAngle()

				currentAngle += (futureAngle - currentAngle) * alpha

				moveable.SetPosition(currentPos)
				moveable.SetAngle(currentAngle)
			}
			// otherwise, don't even try to move it, this object has stopped moving
		}
	}
}

func (m *Manager) GetChannel() byte {
	return 0x04
}

func (m *Manager) HandleFutureData(decoder *gob.Decoder, ttl time.Duration) error {
	var colliderMap = &map[types.ID]collider.Collider{}

	err := decoder.Decode(colliderMap)
	if err != nil {
		fmt.Println("rest", err)
		return err
	}

	m.dataLifeStart = time.Now()
	m.dataLifeEnd = m.dataLifeStart.Add(ttl * time.Millisecond)
	m.colliders = *colliderMap
	return nil
}
