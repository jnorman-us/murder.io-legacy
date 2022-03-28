package worldin

import (
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/types"
)

type Output interface {
	AddByData(types.Channel, packets.Data)
	UpdateByData(types.Channel, packets.Data)
	RemoveByData(types.Channel, packets.Data)
}
