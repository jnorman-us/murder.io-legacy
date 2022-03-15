package ws

import (
	"github.com/josephnormandev/murder/common/communications"
	"github.com/josephnormandev/murder/common/types/action"
	"github.com/josephnormandev/murder/common/types/timestamp"
	"time"
)

type Manager struct {
	timestamp.Timestamp
	communications.Codec

	spawner *Spawner
	// systems         map[byte]*System
	// listeners       map[byte]*Listener
	// futureListeners map[byte]*FutureListener

	receivedFirst bool
	updateQueue   []communications.Clump
	// currentPackets []communications.Packet
}

func NewManager() *Manager {
	var codec = communications.NewCodec()

	return &Manager{
		Codec: *codec,

		// systems:         map[byte]*System{},
		// listeners:       map[byte]*Listener{},
		// futureListeners: map[byte]*FutureListener{},

		receivedFirst: false,
		updateQueue:   make([]communications.Clump, 0),
	}
}

func (m *Manager) SteadyTick(ms time.Duration) {
	var currentTimestamp = m.Tick - 1

	// fmt.Println(len(m.updateQueue), m.Tick)
	for len(m.updateQueue) > 0 {
		var clump = m.updateQueue[0]
		if clump.Timestamp <= currentTimestamp {
			m.EmitToListeners(clump)
			m.updateQueue = m.updateQueue[1:]
		} else if clump.Timestamp == currentTimestamp+1 {
			break
		} else {
			break
		}
	}
	m.TimeTick()
}

func (m *Manager) EncodeSystems() communications.Clump {
	/*var packetArray []communications.Packet

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
	}, nil*/
	return communications.Clump{}
}

func (m *Manager) DecodeForListeners(clump communications.Clump) {
	if !m.receivedFirst {
		m.receivedFirst = true
		m.Tick = clump.Timestamp
	}
	m.updateQueue = append(m.updateQueue, clump)
}

func (m *Manager) EmitToListeners(clump communications.Clump) {
	var spawner = *m.spawner

	for _, p := range clump.Packets {
		var id = p.ID
		var channel = p.Channel
		var act = p.Action
		var datum = p.Data

		if id == 0 {
		} else {
			var class = channel
			switch act {
			case action.Actions.Add:
				spawner.HandleAddition(id, class, datum)
			case action.Actions.Update:
				spawner.HandleUpdate(id, class, datum)
			case action.Actions.Delete:
				spawner.HandleDeletion(id, class, datum)
			}
		}
	}
}
