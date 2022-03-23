package worldin

import (
	"github.com/josephnormandev/murder/client/ws"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
)

type Manager struct {
	packets   *ws.Manager
	listeners map[types.Channel]*ws.Listener
}

func NewManager(pm *ws.Manager, channels []types.Channel) *Manager {
	var listeners = map[types.Channel]*ws.Listener{}
	var manager = &Manager{
		packets: pm,
	}

	var handler = ws.Handler(manager)
	for _, channel := range channels {
		listeners[channel] = pm.CreateListener(channel, &handler)
	}
	return manager
}

func (m *Manager) HandleAdd(c types.Channel, data packets.Data) {
}

func (m *Manager) HandleUpdate(c types.Channel, data packets.Data) {
	switch c {
	}

}

func (m *Manager) HandleDelete(c types.Channel, data packets.Data) {

}
