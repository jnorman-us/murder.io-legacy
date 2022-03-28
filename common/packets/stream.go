package packets

import (
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/types/action"
	"github.com/josephnormandev/murder/common/types/timestamp"
)

type Stream struct {
	timestamp *timestamp.Timestamp
	channel   types.Channel
	catchup   bool

	anonymousID types.ID // if the Data is anonymous, give it this sequence types.ID
	data        map[types.ID]*Data
	additions   chan Addition
	deletions   chan Deletion
}

func newStream(c types.Channel, t *timestamp.Timestamp, catchup bool) *Stream {
	var stream = &Stream{
		channel:   c,
		timestamp: t,
		catchup:   catchup,

		data:      map[types.ID]*Data{},
		additions: make(chan Addition, 100),
		deletions: make(chan Deletion, 100),
	}

	return stream
}

func (s *Stream) CreateAnonymousData(c types.Channel) *Data {
	s.anonymousID++
	var id = s.anonymousID
	return s.CreateData(id, c)
}

func (s *Stream) CreateData(id types.ID, c types.Channel) *Data {
	var data = NewData(id, c, s.timestamp)
	s.additions <- Addition{
		Data:   data,
		Offset: s.timestamp.GetOffsetBytes(),
	}
	return data
}

func (s *Stream) DeleteData(id types.ID) {
	var data = s.data[id]
	s.deletions <- Deletion{
		Data:   data,
		Offset: s.timestamp.GetOffsetBytes(),
	}
}

func (s *Stream) GenerateFullPackets() []Packet {
	var packets []Packet

	if s.catchup {
		for _, data := range s.data {
			packets = append(packets, data.GenerateFullPacket())
		}
	}
	return packets
}

func (s *Stream) GeneratePackets() []Packet {
	var packets []Packet
	for _, data := range s.data {
		if data.Dirty() {
			packets = append(packets, data.GeneratePacket(action.Actions.Update, 0))
			data.CleanDirt()
		}
	}

	for {
		select {
		case addition := <-s.additions:
			var id = addition.Data.ID
			var data = addition.Data
			var offset = addition.Offset
			s.data[id] = data
			data.CleanDirt()
			packets = append(packets, data.GeneratePacket(action.Actions.Add, offset))
		case deletion := <-s.deletions:
			var id = deletion.Data.ID
			var data = deletion.Data
			var offset = deletion.Offset
			s.data[id] = data
			data.CleanDirt()
			packets = append(packets, data.GeneratePacket(action.Actions.Delete, offset))
		default:
			return packets
		}
	}
}
