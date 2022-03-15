package bullet

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/communications/data/schemas"
	"github.com/josephnormandev/murder/common/types"
)

func (b *Bullet) GetClass() types.Channel {
	return Class
}

func (b *Bullet) GetData() data.Data {
	var datum = data.NewData(schemas.BulletSchema)
	datum = b.ColliderData(datum)
	return datum
}

func (b *Bullet) FromData(datum data.Data) {
	datum.ApplySchema(schemas.BulletSchema)
	b.ColliderFromData(datum)
}

const Class types.Channel = 0x83
