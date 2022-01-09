package packet

import (
	"github.com/josephnormandev/murder/common/types"
)

type Packet struct {
	SpawnOrSystem types.SpawnOrSystem
	Identifier    string
	ID            int
	Class         string
	Channel       string
	Data          []byte
}

func (m *Manager) CreatePackets() []Packet {
	var packets []Packet
	for _, s := range m.Spawns {
		var spawn = *s
		var id = spawn.GetID()
		var class = spawn.GetClass()

		m.outputs[class].Reset()
		spawn.GetData(m.encoders[class])
		var outputBytes = m.outputs[class].Bytes()

		packets = append(packets, Packet{
			SpawnOrSystem: types.Spawn(),
			Identifier:    m.Identifier,
			ID:            id,
			Class:         class,
			Data:          outputBytes,
		})
	}

	for _, s := range m.Systems {
		var system = *s
		var channel = system.GetChannel()

		m.outputs[channel].Reset()
		system.GetData(m.encoders[channel])
		var outputBytes = m.outputs[channel].Bytes()
		packets = append(packets, Packet{
			SpawnOrSystem: types.System(),
			Identifier:    m.Identifier,
			Channel:       channel,
			Data:          outputBytes,
		})
	}
	return packets
}
