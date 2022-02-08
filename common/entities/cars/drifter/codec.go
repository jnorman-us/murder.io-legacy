package drifter

import "encoding/gob"

func (d *Drifter) GetClass() string {
	return "d"
}

func (d *Drifter) GetData(e *gob.Encoder) error {
	return e.Encode(d)
}
