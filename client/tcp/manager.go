package tcp

import "github.com/josephnormandev/murder/common/packet"

type Manager struct {
	packet.Codec

	identifier string
	systems    map[string]*System
	listeners  map[string]*Listener
}

func NewManager(identifier string) *Manager {
	var codec = packet.NewCodec()

	return &Manager{
		Codec: *codec,

		identifier: identifier,
		systems:    map[string]*System{},
		listeners:  map[string]*Listener{},
	}
}

func (m *Manager) EncodeSystems() {
	var packetArray []packet.Packet

	for _, s := range m.systems {
		var system = *s
		var channel = system.GetChannel()
		var output = m.Outputs[channel]
		var encoder = m.Encoders[channel]

		output.Reset()
		system.GetData(encoder)
		var outputBytes = output.Bytes()

		packetArray = append(packetArray, packet.Packet{
			Channel: channel,
			Data:    outputBytes,
		})
	}

	m.EncodeOutputs(m.identifier, packetArray)
}

func (m *Manager) DecodeForListeners(ps []packet.Packet) {
	for _, packet := range ps {
		var id = packet.ID
		var channel = packet.Channel
		var data = packet.Data

		var l, ok = m.listeners[channel]
		var input = m.Inputs[channel]
		var decoder = m.Decoders[channel]

		if ok {
			var listener = *l

			input.Reset()
			input.Write(data)
			listener.HandleData(id, decoder)
		}
	}
}
