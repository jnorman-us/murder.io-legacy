package match

import (
	"context"
	"fmt"
	"github.com/josephnormandev/murder/client/ws"
	"github.com/josephnormandev/murder/common/types"
	"syscall/js"
	"time"
)

const steadyTime = time.Millisecond * 1000 / 5

func (m *Manager) Connect(ctx context.Context, hostname string, port int, username types.UserID) error {
	m.wsClient = ws.NewClient(m.packets, hostname, port, username)
	m.Username = username
	err := m.wsClient.Connect(ctx)
	fmt.Println("Stopping connection!")
	return err
}

func (m *Manager) SteadyTick(ctx context.Context) error {
	for range time.Tick(steadyTime) {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping steady tick")
			return ctx.Err()
		default:
			m.packets.SteadyTick(steadyTime)
		}
	}
	return nil
}

func (m *Manager) Update(this js.Value, values []js.Value) interface{} {
	var timeElapsed = m.packets.GetOffsetBytes()
	//var timeTotal = steadyTime

	// fmt.Println(m.packets.Timestamp.Tick, m.packets.Timestamp.GetOffset())

	// here, call packets to release a few packets at a time
	m.packets.Trickle(timeElapsed)
	//m.engine.UpdatePhysics(timeElapsed, timeTotal)

	return nil
}
