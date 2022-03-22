package packets

import (
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/types/timestamp"
)

type Manager struct {
	timestamp.Timestamp
	streams map[types.Channel]*Stream
}

func NewManager() *Manager {
	return &Manager{
		Timestamp: *timestamp.NewTimestamp(),

		streams: map[types.Channel]*Stream{},
	}
}

func (m *Manager) CreateStream(channel types.Channel, catchup bool) *Stream {
	var stream = newStream(channel, &m.Timestamp, catchup)
	m.streams[channel] = stream
	return stream
}

func (m *Manager) MarshalFullPackets() Clump {
	var packets []Packet
	for _, stream := range m.streams {
		var streamPackets = stream.GenerateFullPackets()
		packets = append(packets, streamPackets...)
	}
	return Clump{
		Timestamp: m.Tick,
		Packets:   packets,
	}
}

func (m *Manager) MarshalPackets() Clump {
	var packets []Packet
	for _, stream := range m.streams {
		var streamPackets = stream.GeneratePackets()
		packets = append(packets, streamPackets...)
	}
	return Clump{
		Timestamp: m.Tick,
		Packets:   packets,
	}
}
