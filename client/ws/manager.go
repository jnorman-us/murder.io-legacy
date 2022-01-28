package ws

import (
	"github.com/josephnormandev/murder/common/communications"
)

type Manager struct {
	communications.Codec

	spawner   *Spawner
	systems   map[string]*System
	listeners map[string]*Listener
}

func NewManager() *Manager {
	var codec = communications.NewCodec()

	return &Manager{
		Codec: *codec,

		systems:   map[string]*System{},
		listeners: map[string]*Listener{},
	}
}

func (m *Manager) EncodeSystems() (communications.PacketCollection, error) {
	var packetArray []communications.Packet

	for _, s := range m.systems {
		var system = *s
		var channel = system.GetChannel()

		var encoder = m.BeginEncode(channel)
		err := system.GetData(encoder)
		if err != nil {
			return communications.PacketCollection{}, err
		}
		var outputBytes = m.EndEncode(channel)

		packetArray = append(packetArray, communications.Packet{
			Channel: channel,
			ID:      -1,
			Data:    outputBytes,
		})
	}

	return communications.PacketCollection{
		Timestamp:   0,
		PacketArray: packetArray,
	}, nil
}

func (m *Manager) DecodeForListeners(pc communications.PacketCollection) error {
	for _, p := range pc.PacketArray {
		var id = p.ID
		var channel = p.Channel
		var data = p.Data

		if id == 0 {
			var l, ok = m.listeners[channel]
			if ok {
				decoder, err := m.BeginDecode(channel, data)
				if err != nil {
					return err
				}

				var listener = *l
				err = listener.HandleData(decoder)
				if err != nil {
					return err
				}

				m.EndDecode(channel)
			}
		} else {
			var class = channel

			decoder, err := m.BeginDecode(class, data)
			if err != nil {
				return err
			}

			var spawner = *m.spawner
			err = spawner.HandleSpawn(id, class, decoder)
			if err != nil {
				return err
			}
			m.EndDecode(class)
		}
	}
	return nil
}
