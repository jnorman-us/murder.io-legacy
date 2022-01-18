package packet

import (
	"bytes"
	"encoding/gob"
	"log"
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
	var bytes = c.Outputs[channel].Bytes()
	c.encoderMutexes[channel].Unlock()
	return bytes
}

func (c *Codec) AddDecoder(channel string) {
	var channelInput = new(bytes.Buffer)
	c.Inputs[channel] = channelInput
	c.Decoders[channel] = gob.NewDecoder(channelInput)
	c.decoderMutexes[channel] = &sync.Mutex{}
}

func (c *Codec) BeginDecode(channel string, data []byte) *gob.Decoder {
	c.decoderMutexes[channel].Lock()
	c.Inputs[channel].Reset()
	c.Inputs[channel].Read(data)
	return c.Decoders[channel]
}

func (c *Codec) EndDecode(channel string) {
	c.decoderMutexes[channel].Unlock()
}

func (c *Codec) EncodeOutputs(client string, ps []Packet) []byte {
	c.outputBuffer.Reset()
	var err = c.packetEncoder.Encode(Packets{
		Client:  client,
		Packets: ps,
	})

	if err != nil {
		log.Fatal("encoder error:", err)
	}
	return c.outputBuffer.Bytes()
	// fmt.Println(c.outputBuffer.Len(), c.outputBuffer.Bytes())
}

func (c *Codec) DecodeInputs(data []byte) (string, []Packet) {
	var packets = &Packets{}
	c.inputBuffer.Reset()
	c.inputBuffer.Read(data)

	var err = c.packetDecoder.Decode(packets)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	return packets.Client, packets.Packets
}
