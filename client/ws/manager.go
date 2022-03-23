package ws

import (
	"github.com/josephnormandev/murder/common/codec"
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
	"time"
)

type Manager struct {
	codec.Codec
	packets.Manager
	listeners map[types.Channel]*Listener

	receivedFirst bool
	updateQueue   []packets.Clump
	toRelease     []packets.Packet
	releaseI      int
}

func NewManager() *Manager {
	return &Manager{
		Codec:     *codec.NewCodec(),
		Manager:   *packets.NewManager(),
		listeners: map[types.Channel]*Listener{},

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
			m.updateQueue = m.updateQueue[1:]
			m.release(clump)
		} else {
			break
		}
	}
	m.TimeTick()
}

func (m *Manager) release(clump packets.Clump) {
	for ; m.releaseI < len(m.toRelease); m.releaseI++ {
		var packet = m.toRelease[m.releaseI]
		m.emit(packet, 255)
	}
	m.releaseI = 0
	m.toRelease = clump.Packets
}

func (m *Manager) Trickle(elapsed byte) {
	for ; m.releaseI < len(m.toRelease); m.releaseI++ {
		var packet = m.toRelease[m.releaseI]
		if packet.Offset <= elapsed {
			m.emit(packet, elapsed)
		} else {
			break
		}
	}
	for _, listener := range m.listeners {
		listener.trickle(elapsed)
	}
}

func (m *Manager) emit(packet packets.Packet, elapsed byte) {
	var channel = packet.Channel
	if listener, ok := m.listeners[channel]; ok {
		listener.receive(packet, elapsed)
	}
}

func (m *Manager) CreateListener(c types.Channel, h *Handler) *Listener {
	var listener = newListener(c, h)
	m.listeners[c] = listener
	return listener
}

func (m *Manager) Receive(c packets.Clump) {
	if !m.receivedFirst {
		m.receivedFirst = true
		m.Tick = c.Timestamp
	}
	m.updateQueue = append(m.updateQueue, c)
}
