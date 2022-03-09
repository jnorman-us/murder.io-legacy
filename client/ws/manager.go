package ws

import (
	"github.com/josephnormandev/murder/common/communications"
	"github.com/josephnormandev/murder/common/types"
	"time"
)

type Manager struct {
	types.Tick
	communications.Codec

	spawner         *Spawner
	systems         map[byte]*System
	listeners       map[byte]*Listener
	futureListeners map[byte]*FutureListener

	receivedFirst bool
	updateQueue   []communications.PacketCollection
}

func NewManager() *Manager {
	var codec = communications.NewCodec()

	return &Manager{
		Codec: *codec,

		systems:         map[byte]*System{},
		listeners:       map[byte]*Listener{},
		futureListeners: map[byte]*FutureListener{},

		receivedFirst: false,
		updateQueue:   make([]communications.PacketCollection, 0),
	}
}

func (m *Manager) SteadyTick() error {
	var ms = time.Duration(1000 / 20)
	var currentTimestamp = m.Tick - 2

	// fmt.Println(len(m.updateQueue), m.Tick)
	for len(m.updateQueue) > 0 {
		var pc = m.updateQueue[0]
		if pc.Timestamp <= currentTimestamp {
			err := m.EmitToListeners(pc)
			if err != nil {
				return err
			}
			m.updateQueue = m.updateQueue[1:]
		} else if pc.Timestamp == currentTimestamp+1 {
			err := m.EmitToFutureListeners(pc, ms)
			if err != nil {
				return err
			}
			break
		} else {
			break
		}
	}
	m.Tick++
	return nil
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
	if !m.receivedFirst {
		m.receivedFirst = true
		m.Tick = pc.Timestamp
	}
	m.updateQueue = append(m.updateQueue, pc)
	return nil
}

func (m *Manager) EmitToListeners(pc communications.PacketCollection) error {
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

func (m *Manager) EmitToFutureListeners(pc communications.PacketCollection, ms time.Duration) error {
	for _, p := range pc.PacketArray {
		var id = p.ID
		var channel = p.Channel
		var data = p.Data

		if id == 0 {
			var l, ok = m.futureListeners[channel]
			if ok {
				decoder, err := m.BeginDecode(channel, data)
				if err != nil {
					return err
				}

				var listener = *l
				err = listener.HandleFutureData(decoder, ms)
				if err != nil {
					return err
				}

				m.EndDecode(channel)
			}
		}
	}
	return nil
}
