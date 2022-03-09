package match

import (
	"context"
	"fmt"
	"github.com/josephnormandev/murder/client/ws"
	"github.com/josephnormandev/murder/common/types"
	"time"
)

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
			err := m.packets.SteadyTick()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
