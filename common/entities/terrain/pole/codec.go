package pole

import "encoding/gob"

func (p *Pole) GetClass() string {
	return "p"
}

func (p *Pole) GetData(e *gob.Encoder) error {
	return e.Encode(p)
}
