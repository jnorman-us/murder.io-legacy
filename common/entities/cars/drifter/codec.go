package drifter

import (
	"encoding/gob"
)

func (d *Drifter) GetClass() byte {
	return Class
}

func (d *Drifter) GetData(e *gob.Encoder) error {
	return e.Encode(d)
}

const Class byte = 0x81
