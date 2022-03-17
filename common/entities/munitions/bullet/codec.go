package bullet

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/communications/data/schemas"
	"github.com/josephnormandev/murder/common/types"
)

func (b *Bullet) GetClass() types.Channel {
	return Class
}

func (b *Bullet) GetStartData() data.Data {
	var datum = data.NewData(schemas.BulletStartSchema)
	datum.SetFloat("X", b.Position.X)
	datum.SetFloat("Y", b.Position.Y)
	datum.SetFloat("Angle", b.Angle)
	return datum
}

func (b *Bullet) GetData() data.Data {
	var datum = data.NewData(schemas.BulletSchema)
	return datum
}

func (b *Bullet) FromData(datum data.Data) {
}

func (b *Bullet) FromStartData(datum data.Data) {
	datum.ApplySchema(schemas.BulletStartSchema)
	var x = datum.GetFloat("X")
	var y = datum.GetFloat("Y")
	var angle = datum.GetFloat("Angle")
	b.SetPosition(types.NewVector(x, y))
	b.SetAngle(angle)
}

const Class types.Channel = 0x83
