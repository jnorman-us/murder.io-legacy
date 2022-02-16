package dimetrodon

import "encoding/gob"

func (d *Dimetrodon) GetClass() byte {
	return Class
}

func (d *Dimetrodon) GetData(e *gob.Encoder) error {
	return e.Encode(d)
}

const Class byte = 0x84
