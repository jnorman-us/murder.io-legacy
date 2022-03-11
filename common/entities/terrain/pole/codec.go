package pole

import "encoding/gob"

func (p *Pole) GetClass() byte {
	return 0x82
}

func (p *Pole) GetData(e *gob.Encoder) error {
	return e.Encode(p.State)
}

const Class byte = 0x82
