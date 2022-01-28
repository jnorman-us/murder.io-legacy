package communications

import "github.com/josephnormandev/murder/common/types"

type PacketCollection struct {
	Timestamp   types.Tick
	PacketArray []Packet
}
