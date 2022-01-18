package packet

type Packets struct {
	Client  string
	Packets []Packet
}

type Packet struct {
	Channel string // decoder channel
	ID      int    // extra identifier in the packet
	Data    []byte // data to be decoded
}

/*
func (m *Manager) CreatePackets() []Packet {
	var packets []Packet
	for _, s := range m.Spawns {
		var spawn = *s
		var id = spawn.GetID()
		var class = spawn.GetClass()

		m.outputs[class].Reset()
		spawn.GetData(m.encoders[class])
		var outputBytes = m.outputs[class].Bytes()

		packets = append(packets, Packet{
			SpawnOrSystem: types.Spawn(),
			ID:            id,
			Class:         class,
			Data:          outputBytes,
		})
	}

	for _, s := range m.Systems {
		var system = *s
		var channel = system.GetChannel()

		m.outputs[channel].Reset()
		system.GetData(m.encoders[channel])
		var outputBytes = m.outputs[channel].Bytes()
		packets = append(packets, Packet{
			SpawnOrSystem: types.System(),
			Identifier:    m.Identifier,
			Channel:       channel,
			Data:          outputBytes,
		})
	}
	return packets
}
*/
