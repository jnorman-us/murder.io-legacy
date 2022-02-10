package communications

import "github.com/josephnormandev/murder/common/types"

type Packet struct {
	Channel byte     // decoder channel
	ID      types.ID // extra identifier in the communications
	Data    []byte   // data to be decoded
}
