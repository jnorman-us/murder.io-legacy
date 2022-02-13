package bullet

import "encoding/gob"

func (b *Bullet) GetClass() byte {
	return Class
}

func (b *Bullet) GetData(e *gob.Encoder) error {
	return e.Encode(b)
}

const Class byte = 0x83
