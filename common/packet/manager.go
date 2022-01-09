package packet

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Manager struct {
	Identifier string
	Systems    map[string]*System     // the systems that want to send data
	Listeners  map[string][]*Listener // the systems that want to hear data

	Spawns  map[int]*Spawn // the entities that change on server
	Spawner *Spawner       // the syncer that updates client on received spawns

	decoders map[string]*gob.Decoder
	inputs   map[string]*bytes.Buffer
	encoders map[string]*gob.Encoder
	outputs  map[string]*bytes.Buffer

	packetEncoder *gob.Encoder
	outputBuffer  *bytes.Buffer
	packetDecoder *gob.Decoder
	inputBuffer   *bytes.Buffer
}

func NewManager(id string) *Manager {
	var outputBuffer = new(bytes.Buffer)
	var inputBuffer = new(bytes.Buffer)

	return &Manager{
		Identifier: id,
		Systems:    map[string]*System{},
		Listeners:  map[string][]*Listener{},

		Spawns: map[int]*Spawn{},

		decoders: map[string]*gob.Decoder{},
		inputs:   map[string]*bytes.Buffer{},
		encoders: map[string]*gob.Encoder{},
		outputs:  map[string]*bytes.Buffer{},

		packetDecoder: gob.NewDecoder(inputBuffer),
		packetEncoder: gob.NewEncoder(outputBuffer),
		outputBuffer:  outputBuffer,
		inputBuffer:   inputBuffer,
	}
}

func (m *Manager) EncodeOutputs() {
	var packets = m.CreatePackets()
	m.outputBuffer.Reset()
	var err = m.packetEncoder.Encode(packets)

	if err != nil {
		log.Fatal("encoder error:", err)
	}
	fmt.Println(m.outputBuffer.Len(), m.outputBuffer.Bytes())
}

func (m *Manager) DecodeInputs() {
	var packets = &[]Packet{}
	var err = m.packetDecoder.Decode(packets)

	if err != nil {
		log.Fatal("decode error:", err)
	}

	for _, packet := range *packets {
		if packet.SpawnOrSystem.IsSpawn() {
			var class = packet.Class
			var id = packet.ID
			var data = packet.Data

			var input, ok = m.inputs[class]
			var decoder = m.decoders[class]
			if !ok {
				input = new(bytes.Buffer)
				decoder = gob.NewDecoder(input)
				m.inputs[class] = input
				m.decoders[class] = decoder
			}

			var spawner = *m.Spawner
			input.Reset()
			input.Write(data)
			spawner.AddSpawn(id, class, decoder)
		} else if packet.SpawnOrSystem.IsSystem() {
			var channel = packet.Channel
			var data = packet.Data

			var input, ok = m.inputs[channel]
			var decoder = m.decoders[channel]
			if !ok {
				input = new(bytes.Buffer)
				decoder = gob.NewDecoder(input)
				m.inputs[channel] = input
				m.decoders[channel] = decoder
			}

			for _, l := range m.Listeners[channel] {
				var listener = *l
				input.Reset()
				input.Write(data)
				listener.HandleData(decoder)
			}
		}
	}
}

func (m *Manager) CopyOver() {
	m.inputBuffer.Reset()
	m.inputBuffer.ReadFrom(m.outputBuffer)
}
