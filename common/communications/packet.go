package communications

type Packet struct {
	Channel string // decoder channel
	ID      int    // extra identifier in the communications
	Data    []byte // data to be decoded
}
