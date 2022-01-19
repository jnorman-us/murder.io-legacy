package ws

import "github.com/josephnormandev/murder/common/packet"

type Manager struct {
	packet.Codec

	spawner   *Spawner
	systems   map[string]*System
	listeners map[string]*Listener
}

func NewManager() *Manager {
	var codec = packet.NewCodec()

	return &Manager{
		Codec: *codec,

		systems:   map[string]*System{},
		listeners: map[string]*Listener{},
	}
}

func (m *Manager) EncodeSystems() ([]packet.Packet, error) {
	var packetArray []packet.Packet

	for _, s := range m.systems {
		var system = *s
		var channel = system.GetChannel()

		var encoder = m.BeginEncode(channel)
		err := system.GetData(encoder)
		if err != nil {
			return []packet.Packet{}, err
		}
		var outputBytes = m.EndEncode(channel)

		packetArray = append(packetArray, packet.Packet{
			Channel: channel,
			ID:      -1,
			Data:    outputBytes,
		})
	}

	return packetArray, nil
}

func (m *Manager) DecodeForListeners(ps []packet.Packet) error {
	for _, p := range ps {
		var id = p.ID
		var channel = p.Channel
		var data = p.Data

		if id == -1 {
			var l, ok = m.listeners[channel]
			if ok {
				decoder, err := m.BeginDecode(channel, data)
				if err != nil {
					return err
				}
				var listener = *l
				err = listener.HandleData(id, decoder)
				if err != nil {
					return err
				}
				m.EndDecode(channel)
			}
		} else {
			var class = channel
			var spawner = *m.spawner

			decoder, err := m.BeginDecode(class, data)
			if err != nil {
				return err
			}
			spawner.HandleSpawn(id, class, decoder)
			m.EndDecode(class)
		}
	}
	return nil
}
