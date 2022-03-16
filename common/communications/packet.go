package communications

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/types"
	"github.com/josephnormandev/murder/common/types/action"
	"time"
)

type Packet struct {
	ID      types.ID
	Channel types.Channel // decoder channel
	Action  action.Action
	Offset  byte
	Data    data.Data
}

func NewPacket(channel types.Channel, action action.Action, offset time.Duration, dat data.Data) Packet {
	return Packet{
		Channel: channel,
		Action:  action,
		Offset:  byte(offset / 1000000),
		Data:    dat,
	}
}

func NewSpawnPacket(id types.ID, class types.Channel, action action.Action, offset time.Duration, datum data.Data) Packet {
	return Packet{
		ID:      id,
		Channel: class,
		Action:  action,
		Offset:  byte(offset / time.Millisecond),
		Data:    datum,
	}
}

func (p *Packet) GetOffset() time.Duration {
	return time.Duration(p.Offset) * time.Millisecond
}
