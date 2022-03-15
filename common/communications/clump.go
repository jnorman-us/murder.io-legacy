package communications

import (
	"github.com/josephnormandev/murder/common/types/timestamp"
)

type Clump struct {
	Timestamp timestamp.Tick
	Packets   []Packet
}
