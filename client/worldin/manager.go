package worldin

import (
	"github.com/josephnormandev/murder/client/ws"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/packets/schemas"
	"github.com/josephnormandev/murder/common/types"
)

type Manager struct {
	packets   *ws.Manager
	listeners map[types.Channel]*ws.Listener
	output    *Output
}

func NewManager(pm *ws.Manager, output *Output) *Manager {
	var channels = []types.Channel{
		schemas.BulletSchema.Channel(),
		schemas.DimetrodonSchema.Channel(),
		schemas.PoleSchema.Channel(),
	}
	var listeners = map[types.Channel]*ws.Listener{}
	var manager = &Manager{
		packets: pm,
		output:  output,
	}

	var handler = ws.Handler(manager)
	for _, channel := range channels {
		listeners[channel] = pm.CreateListener(channel, &handler)
	}
	return manager
}

func (m *Manager) HandleAdd(c types.Channel, data packets.Data) {
	(*m.output).AddByData(c, data)
}

func (m *Manager) HandleUpdate(c types.Channel, data packets.Data) {
	(*m.output).UpdateByData(c, data)
}

func (m *Manager) HandleDelete(c types.Channel, data packets.Data) {
	(*m.output).RemoveByData(c, data)
}
