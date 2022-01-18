package websocket

import "github.com/josephnormandev/murder/common/packet"

type Manager struct {
	packet.Codec

	identifier string
	spawner    *Spawner
	systems    map[string]*System
	listeners  map[string]*Listener
}

func NewManager(identifier string) *Manager {
	var codec = packet.NewCodec()

	return &Manager{
		Codec: *codec,

		identifier: identifier,
		systems:    map[string]*System{},
		listeners:  map[string]*Listener{},
	}
}

func (m *Manager) EncodeSystems() {
	var packetArray []packet.Packet

	for _, s := range m.systems {
		var system = *s
		var channel = system.GetChannel()

		var encoder = m.BeginEncode(channel)
		system.GetData(encoder)
		var outputBytes = m.EndEncode(channel)

		packetArray = append(packetArray, packet.Packet{
			Channel: channel,
			Data:    outputBytes,
		})
	}

	m.EncodeOutputs(m.identifier, packetArray)
}

func (m *Manager) DecodeForListeners(ps []packet.Packet) {
	for _, packet := range ps {
		var id = packet.ID
		var channel = packet.Channel
		var data = packet.Data

		if id == -1 {
			var l, ok = m.listeners[channel]
			if ok {
				var decoder = m.BeginDecode(channel, data)
				var listener = *l
				listener.HandleData(id, decoder)
				m.EndDecode(channel)
			}
		} else {
			// pass it to the spawner, which will take this packet
			// as some kinda entity
			var class = channel
			var spawner = *m.spawner

			var decoder = m.BeginDecode(class, data)
			spawner.HandleSpawn(id, class, decoder)
			m.EndDecode(class)
		}
	}
}
