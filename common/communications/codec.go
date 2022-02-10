package communications

import (
	"bytes"
	"encoding/gob"
)

type Codec struct {
	Decoders map[byte]*gob.Decoder
	Inputs   map[byte]*bytes.Buffer
	Encoders map[byte]*gob.Encoder
	Outputs  map[byte]*bytes.Buffer

	packetEncoder *gob.Encoder
	outputBuffer  *bytes.Buffer
	packetDecoder *gob.Decoder
	inputBuffer   *bytes.Buffer
}

func NewCodec() *Codec {
	var outputBuffer = new(bytes.Buffer)
	var inputBuffer = new(bytes.Buffer)

	return &Codec{
		Decoders: map[byte]*gob.Decoder{},
		Inputs:   map[byte]*bytes.Buffer{},
		Encoders: map[byte]*gob.Encoder{},
		Outputs:  map[byte]*bytes.Buffer{},

		packetDecoder: gob.NewDecoder(inputBuffer),
		packetEncoder: gob.NewEncoder(outputBuffer),
		outputBuffer:  outputBuffer,
		inputBuffer:   inputBuffer,
	}
}

func (c *Codec) AddEncoder(channel byte) {
	var channelOutput = new(bytes.Buffer)
	c.Outputs[channel] = channelOutput
	c.Encoders[channel] = gob.NewEncoder(channelOutput)
}

func (c *Codec) BeginEncode(channel byte) *gob.Encoder {
	c.Outputs[channel].Reset()
	return c.Encoders[channel]
}

func (c *Codec) EndEncode(channel byte) []byte {
	var byteArray = make([]byte, c.Outputs[channel].Len())
	copy(byteArray, c.Outputs[channel].Bytes())
	return byteArray
}

func (c *Codec) AddDecoder(channel byte) {
	var channelInput = new(bytes.Buffer)
	c.Inputs[channel] = channelInput
	c.Decoders[channel] = gob.NewDecoder(channelInput)
}

func (c *Codec) BeginDecode(channel byte, data []byte) (*gob.Decoder, error) {
	c.Inputs[channel].Reset()
	_, err := c.Inputs[channel].Write(data)

	if err != nil {
		return nil, err
	}
	return c.Decoders[channel], nil
}

func (c *Codec) EndDecode(channel byte) {
}

func (c *Codec) EncodeOutputs(pc PacketCollection) ([]byte, error) {
	c.outputBuffer.Reset()
	var err = c.packetEncoder.Encode(pc)
	if err != nil {
		return []byte{}, err
	}

	var byteArray = make([]byte, c.outputBuffer.Len())
	copy(byteArray, c.outputBuffer.Bytes())
	return byteArray, nil
}

func (c *Codec) DecodeInputs(data []byte) (PacketCollection, error) {
	var packets = &PacketCollection{}
	c.inputBuffer.Reset()
	_, err := c.inputBuffer.Write(data)
	if err != nil {
		return PacketCollection{}, err
	}

	err = c.packetDecoder.Decode(packets)
	if err != nil {
		return PacketCollection{}, nil
	}
	return *packets, nil
}
