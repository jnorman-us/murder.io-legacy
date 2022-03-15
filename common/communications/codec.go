package communications

import (
	"bytes"
	"encoding/gob"
)

type Codec struct {
	encoder *gob.Encoder
	decoder *gob.Decoder

	input  *bytes.Buffer
	output *bytes.Buffer
}

func NewCodec() *Codec {
	var input = new(bytes.Buffer)
	var output = new(bytes.Buffer)

	return &Codec{
		decoder: gob.NewDecoder(input),
		encoder: gob.NewEncoder(output),

		input:  input,
		output: output,
	}
}

func (c *Codec) EncodeOutputs(pc Clump) ([]byte, error) {
	c.output.Reset()
	var err = c.encoder.Encode(pc)
	if err != nil {
		return []byte{}, err
	}

	var outputs = make([]byte, c.output.Len())
	copy(outputs, c.output.Bytes())
	return outputs, nil
}

func (c *Codec) DecodeInputs(data []byte) (Clump, error) {
	var clump = &Clump{}
	c.input.Reset()
	_, err := c.input.Write(data)
	if err != nil {
		return Clump{}, err
	}

	err = c.decoder.Decode(clump)
	if err != nil {
		return Clump{}, nil
	}
	return *clump, nil
}
