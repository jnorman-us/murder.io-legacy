package engine

import (
	"encoding/gob"
	"fmt"
	"github.com/josephnormandev/murder/common/collisions/collider"
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

func (m *Manager) UpdatePhysics(ms float64) {
	var currentTime = time.Now()
	var timeElapsed = currentTime.Sub(m.dataLifeStart)
	var timeTotal = m.dataLifeEnd.Sub(m.dataLifeStart)

	var alpha = 0.0
	if timeTotal != 0 {
		alpha = float64(timeElapsed) / float64(timeTotal)
		alpha = math.Min(1, alpha)
	}

	for id, mo := range m.moveables {
		var moveable = *mo
		// var collider, ok = m.colliders[id]
		var collider, ok1 = m.colliders[id]
		if ok1 {
			var currentPos = moveable.GetPosition()
			var futurePos = collider.GetPosition()

			currentPos.Interpolate(futurePos, alpha)

			var currentAngle = moveable.GetAngle()
			var futureAngle = collider.GetAngle()

			currentAngle += (futureAngle - currentAngle) * alpha

			moveable.SetPosition(currentPos)
			moveable.SetAngle(currentAngle)
		} else {
			moveable.UpdatePosition(ms / (1000 / 60))
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
