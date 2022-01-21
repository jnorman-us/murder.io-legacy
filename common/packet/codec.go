package packet

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
)

type Codec struct {
	Decoders map[string]*gob.Decoder
	Inputs   map[string]*bytes.Buffer
	Encoders map[string]*gob.Encoder
	Outputs  map[string]*bytes.Buffer

	encoderMutexes map[string]*sync.Mutex
	decoderMutexes map[string]*sync.Mutex

	packetEncoder *gob.Encoder
	outputBuffer  *bytes.Buffer
	packetDecoder *gob.Decoder
	inputBuffer   *bytes.Buffer

	encoderMutex *sync.Mutex
	decoderMutex *sync.Mutex
}

func NewCodec() *Codec {
	var outputBuffer = new(bytes.Buffer)
	var inputBuffer = new(bytes.Buffer)

	return &Codec{
		Decoders: map[string]*gob.Decoder{},
		Inputs:   map[string]*bytes.Buffer{},
		Encoders: map[string]*gob.Encoder{},
		Outputs:  map[string]*bytes.Buffer{},

		encoderMutexes: map[string]*sync.Mutex{},
		decoderMutexes: map[string]*sync.Mutex{},

		packetDecoder: gob.NewDecoder(inputBuffer),
		packetEncoder: gob.NewEncoder(outputBuffer),
		outputBuffer:  outputBuffer,
		inputBuffer:   inputBuffer,

		encoderMutex: &sync.Mutex{},
		decoderMutex: &sync.Mutex{},
	}
}

func (c *Codec) AddEncoder(channel string) {
	var channelOutput = new(bytes.Buffer)
	c.Outputs[channel] = channelOutput
	c.Encoders[channel] = gob.NewEncoder(channelOutput)
	c.encoderMutexes[channel] = &sync.Mutex{}
}

func (c *Codec) BeginEncode(channel string) *gob.Encoder {
	c.encoderMutexes[channel].Lock()
	c.Outputs[channel].Reset()
	return c.Encoders[channel]
}

func (c *Codec) EndEncode(channel string) []byte {
	var byteArray = c.Outputs[channel].Bytes()
	c.encoderMutexes[channel].Unlock()
	return byteArray
}

func (c *Codec) AddDecoder(channel string) {
	var channelInput = new(bytes.Buffer)
	c.Inputs[channel] = channelInput
	c.Decoders[channel] = gob.NewDecoder(channelInput)
	c.decoderMutexes[channel] = &sync.Mutex{}
}

func (c *Codec) BeginDecode(channel string, data []byte) (*gob.Decoder, error) {
	c.decoderMutexes[channel].Lock()
	c.Inputs[channel].Reset()
	_, err := c.Inputs[channel].Write(data)

	if err != nil {
		return nil, err
	}
	return c.Decoders[channel], nil
}

func (c *Codec) EndDecode(channel string) {
	c.decoderMutexes[channel].Unlock()
}

func (c *Codec) EncodeOutputs(ps []Packet) ([]byte, error) {
	fmt.Println("Packet Output!", ps)
	c.encoderMutex.Lock()
	defer c.encoderMutex.Unlock()

	c.outputBuffer.Reset()
	var err = c.packetEncoder.Encode(ps)
	if err != nil {
		return []byte{}, err
	}

	var byteArray = c.outputBuffer.Bytes()
	return byteArray, nil
}

func (c *Codec) DecodeInputs(data []byte) ([]Packet, error) {
	c.decoderMutex.Lock()
	defer c.decoderMutex.Unlock()

	var packets = &[]Packet{}
	c.inputBuffer.Reset()
	_, err := c.inputBuffer.Write(data)
	if err != nil {
		fmt.Printf("Error decoding! %v\n", err)
		return []Packet{}, err
	}

	err = c.packetDecoder.Decode(packets)
	if err != nil {
		return []Packet{}, nil
	}
	fmt.Println("Packet Input!", packets)
	return *packets, nil
}
