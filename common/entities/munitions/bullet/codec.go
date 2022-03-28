package bullet

import (
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/packets/schemas"
	"github.com/josephnormandev/murder/common/types"
)

func (b *Bullet) GetClass() types.Channel {
	return Class
}

func (b *Bullet) PopulateData(data *packets.Data) {
	data.SetFloat("X", b.Position.X)
	data.SetFloat("Y", b.Position.Y)
	data.SetFloat("Angle", b.Angle)
}

func (b *Bullet) FromData(data packets.Data) {
	var position = types.NewVector(data.GetFloat("X"), data.GetFloat("Y"))
	var angle = data.GetFloat("Angle")
	b.SetPosition(position)
	b.SetAngle(angle)
}

var Class = schemas.BulletSchema.Channel()
