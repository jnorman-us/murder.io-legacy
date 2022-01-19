package packet

type Packet struct {
	Channel string // decoder channel
	ID      int    // extra identifier in the packet
	Data    []byte // data to be decoded
}
