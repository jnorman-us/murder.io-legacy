package packet

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Codec struct {
	Decoders map[string]*gob.Decoder
	Inputs   map[string]*bytes.Buffer
	Encoders map[string]*gob.Encoder
	Outputs  map[string]*bytes.Buffer

	packetEncoder *gob.Encoder
	outputBuffer  *bytes.Buffer
	packetDecoder *gob.Decoder
	inputBuffer   *bytes.Buffer
}

func NewCodec() *Codec {
	var outputBuffer = new(bytes.Buffer)
	var inputBuffer = new(bytes.Buffer)

	return &Codec{
		Decoders: map[string]*gob.Decoder{},
		Inputs:   map[string]*bytes.Buffer{},
		Encoders: map[string]*gob.Encoder{},
		Outputs:  map[string]*bytes.Buffer{},

		packetDecoder: gob.NewDecoder(inputBuffer),
		packetEncoder: gob.NewEncoder(outputBuffer),
		outputBuffer:  outputBuffer,
		inputBuffer:   inputBuffer,
	}
}

func (c *Codec) AddEncoder(channel string) {
	var channelOutput = new(bytes.Buffer)
	c.Outputs[channel] = channelOutput
	c.Encoders[channel] = gob.NewEncoder(channelOutput)
}

func (c *Codec) AddDecoder(channel string) {
	var channelInput = new(bytes.Buffer)
	c.Inputs[channel] = channelInput
	c.Decoders[channel] = gob.NewDecoder(channelInput)
}

func (c *Codec) EncodeOutputs(client string, ps []Packet) {
	c.outputBuffer.Reset()
	var err = c.packetEncoder.Encode(Packets{
		Client:  client,
		Packets: ps,
	})

	if err != nil {
		log.Fatal("encoder error:", err)
	}
	// fmt.Println(c.outputBuffer.Len(), c.outputBuffer.Bytes())
}

func (c *Codec) DecodeInputs() (string, []Packet) {
	var packets = &Packets{}
	var err = c.packetDecoder.Decode(packets)

	if err != nil {
		log.Fatal("decode error:", err)
	}

	return packets.Client, packets.Packets
}

func (c *Codec) GetOutputBuffer() *bytes.Buffer {
	return c.outputBuffer
}

func (c *Codec) GetInputBuffer() *bytes.Buffer {
	return c.inputBuffer
}
