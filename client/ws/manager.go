package ws

import (
	"github.com/josephnormandev/murder/common/communications"
	"github.com/josephnormandev/murder/common/types/action"
	"github.com/josephnormandev/murder/common/types/timestamp"
	"sort"
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
	toRelease     []communications.Packet
	releaseI      int
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
			var packets = clump.Packets
			m.updateQueue = m.updateQueue[1:]
			m.setRelease(packets)
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

func (m *Manager) setRelease(packets []communications.Packet) {
	sort.Slice(packets, func(i, j int) bool {
		return packets[i].Offset < packets[j].Offset
	})
	for _, packet := range m.toRelease {
		m.emit(packet)
	}
	m.toRelease = packets
	m.releaseI = 0
}

func (m *Manager) TrickleEmit(elapsed time.Duration) {
	for ; m.releaseI < len(m.toRelease); m.releaseI++ {
		var p = m.toRelease[m.releaseI]

		if p.GetOffset() < elapsed {
			m.emit(p)
		}
	}
}

func (m *Manager) emit(p communications.Packet) {
	var spawner = *m.spawner

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
