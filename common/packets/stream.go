package packets

import (
	"fmt"
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

func (s *Stream) CreateAnonymousData(sc Schema) *Data {
	s.anonymousID++
	var id = s.anonymousID
	return s.CreateData(id, sc)
}

func (s *Stream) CreateData(id types.ID, sc Schema) *Data {
	var data = newData(id, sc, s.timestamp)
	fmt.Println("ADDING")
	s.additions <- Addition{
		data:   data,
		offset: s.timestamp.GetOffsetBytes(),
	}
	fmt.Println("DONE ADDING")
	return data
}

func (s *Stream) DeleteData(id types.ID) {
	var data = s.data[id]
	s.deletions <- Deletion{
		data:   data,
		offset: s.timestamp.GetOffsetBytes(),
	}
}

func (s *Stream) GenerateFullPackets() []Packet {
	var packets []Packet

	if s.catchup {
		var channel = s.channel
		for _, data := range s.data {
			packets = append(packets, data.GenerateFullPacket(channel))
		}
	}
	return packets
}

func (s *Stream) GeneratePackets() []Packet {
	var packets []Packet
	var channel = s.channel

	for _, data := range s.data {
		if data.Dirty() {
			packets = append(packets, data.GeneratePacket(channel, action.Actions.Update, 0))
			data.CleanDirt()
		}
	}

	for {
		select {
		case addition := <-s.additions:
			var id = addition.data.id
			var data = addition.data
			var offset = addition.offset
			s.data[id] = data
			data.CleanDirt()
			packets = append(packets, data.GeneratePacket(channel, action.Actions.Add, offset))
		case deletion := <-s.deletions:
			var id = deletion.data.id
			var data = deletion.data
			var offset = deletion.offset
			s.data[id] = data
			data.CleanDirt()
			packets = append(packets, data.GeneratePacket(channel, action.Actions.Delete, offset))
		default:
			fmt.Println(len(s.additions), len(s.deletions))
			return packets
		}
	}
}
