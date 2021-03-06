package pole

import (
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/packets/schemas"
	"github.com/josephnormandev/murder/common/types"
)

func (p *Pole) GetClass() types.Channel {
	return Class
}

func (p *Pole) PopulateData(data *packets.Data) {
	data.SetFloat("X", p.Position.X)
	data.SetFloat("Y", p.Position.Y)
	data.SetFloat("Angle", p.Angle)
}

func (p *Pole) FromData(data packets.Data) {
	var position = types.NewVector(data.GetFloat("X"), data.GetFloat("Y"))
	var angle = data.GetFloat("Angle")
	p.SetPosition(position)
	p.SetAngle(angle)
}

var Class = schemas.PoleSchema.Channel()
