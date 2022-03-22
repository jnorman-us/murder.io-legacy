package ws

import (
	"github.com/josephnormandev/murder/common/codec"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/types/action"
	"github.com/josephnormandev/murder/common/types/timestamp"
	"sort"
	"time"
)

type Manager struct {
	timestamp.Timestamp
	codec.Codec

	spawner *Spawner
	// systems         map[byte]*System
	listeners       map[types.Channel]*Listener
	futureListeners map[types.Channel]*FutureListener

	receivedFirst bool
	updateQueue   []packets.Clump
	toRelease     []packets.Packet
	releaseI      int
}

func NewManager() *Manager {
	var codec = codec.NewCodec()

	return &Manager{
		Codec: *codec,

		// systems:         map[byte]*System{},
		listeners:       map[types.Channel]*Listener{},
		futureListeners: map[types.Channel]*FutureListener{},

		receivedFirst: false,
		updateQueue:   make([]packets.Clump, 0),
	}
}

func (m *Manager) SteadyTick(ms time.Duration) {
	var currentTimestamp = m.Tick - 2

	// fmt.Println(len(m.updateQueue), m.Tick)
	for len(m.updateQueue) > 0 {
		var clump = m.updateQueue[0]
		if clump.Timestamp <= currentTimestamp {
			var packets = clump.Packets
			m.updateQueue = m.updateQueue[1:]
			m.setRelease(packets)
		} else if clump.Timestamp == currentTimestamp+1 {
			var packets = clump.Packets
			m.EmitFutures(packets)
			break
		} else {
			break
		}
	}
	m.TimeTick()
}

func (m *Manager) EncodeSystems() packets.Clump {
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
	return packets.Clump{}
}

func (m *Manager) DecodeForListeners(clump packets.Clump) {
	if !m.receivedFirst {
		m.receivedFirst = true
		m.Tick = clump.Timestamp
	}
	m.updateQueue = append(m.updateQueue, clump)
}

func (m *Manager) setRelease(packets []packets.Packet) {
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

func (m *Manager) emit(p packets.Packet) {
	var spawner = *m.spawner

	var id = p.ID
	var channel = p.Channel
	var act = p.Action
	if id == 0 {
		var datums = p.Datas
		if l, ok := m.listeners[channel]; ok {
			(*l).HandleData(datums)
		}
	} else {
		var class = channel
		var datum = p.Data
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

func (m *Manager) EmitFutures(packets []packets.Packet) {
	for _, p := range packets {
		var channel = p.Channel
		if l, ok := m.futureListeners[channel]; ok {
			(*l).HandleFutureData(p.Datas)
		}
	}
}
