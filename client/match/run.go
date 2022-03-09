package match

import (
	"fmt"
	"github.com/josephnormandev/murder/client/ws"
	"github.com/josephnormandev/murder/common/types"
	"syscall/js"
	"time"
)

func (m *Manager) Connect(this js.Value, values []js.Value) interface{} {
	var hostname = values[0].String()
	var port = values[1].Int()
	var username = types.UserID(values[2].String())

	m.RunGroup.Go(func() error {
		m.wsClient = ws.NewClient(m.packets, hostname, port, username)
		err := m.wsClient.Connect(m.RunContext)
		fmt.Println("Stopping connection!")
		return err
	})
	return nil
}

func (m *Manager) UpdateTick() {
	m.RunGroup.Go(func() error {
		for range time.Tick(updateTime) {
			select {
			case <-m.RunContext.Done():
				fmt.Println("Stopping update tick")
				return m.RunContext.Err()
			default:
				m.engine.UpdatePhysics(updateTime)
			}
		}
		return nil
	})
}

func (m *Manager) SteadyTick() {
	m.RunGroup.Go(func() error {
		for range time.Tick(steadyTime) {
			select {
			case <-m.RunContext.Done():
				fmt.Println("Stopping steady tick")
				return m.RunContext.Err()
			default:
				err := m.packets.SteadyTick()
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}
