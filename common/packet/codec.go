package packet

import (
	"bytes"
	"encoding/gob"
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

func (c *Codec) BeginEncode(channel string) *gob.Encoder {
	c.Outputs[channel].Reset()
	return c.Encoders[channel]
}

func (c *Codec) EndEncode(channel string) []byte {
	var byteArray = make([]byte, c.Outputs[channel].Len())
	copy(byteArray, c.Outputs[channel].Bytes())
	return byteArray
}

func (c *Codec) AddDecoder(channel string) {
	var channelInput = new(bytes.Buffer)
	c.Inputs[channel] = channelInput
	c.Decoders[channel] = gob.NewDecoder(channelInput)
}

func (c *Codec) BeginDecode(channel string, data []byte) (*gob.Decoder, error) {
	c.Inputs[channel].Reset()
	_, err := c.Inputs[channel].Write(data)

	if err != nil {
		return nil, err
	}
	return c.Decoders[channel], nil
}

func (c *Codec) EndDecode(channel string) {
}

func (c *Codec) EncodeOutputs(ps []Packet) ([]byte, error) {
	c.outputBuffer.Reset()
	var err = c.packetEncoder.Encode(ps)
	if err != nil {
		return []byte{}, err
	}

	var byteArray = make([]byte, c.outputBuffer.Len())
	copy(byteArray, c.outputBuffer.Bytes())
	return byteArray, nil
}

func (c *Codec) DecodeInputs(data []byte) ([]Packet, error) {
	var packets = &[]Packet{}
	c.inputBuffer.Reset()
	_, err := c.inputBuffer.Write(data)
	if err != nil {
		return []Packet{}, err
	}

	err = c.packetDecoder.Decode(packets)
	if err != nil {
		return []Packet{}, nil
	}
	return *packets, nil
}
