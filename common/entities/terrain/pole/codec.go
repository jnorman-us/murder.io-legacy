package pole

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/communications/data/schemas"
	"github.com/josephnormandev/murder/common/types"
)

func (p *Pole) GetClass() types.Channel {
	return 0x82
}

func (p *Pole) GetStartData() data.Data {
	var datum = data.NewData(schemas.PoleStartSchema)
	datum.SetFloat("X", p.Position.X)
	datum.SetFloat("Y", p.Position.Y)
	datum.SetFloat("Angle", p.Angle)
	return datum
}

func (p *Pole) GetData() data.Data {
	var datum = data.NewData(schemas.PoleSchema)
	return datum
}

func (p *Pole) FromData(datum data.Data) {
}

func (p *Pole) FromStartData(datum data.Data) {
	datum.ApplySchema(schemas.PoleStartSchema)
	var x = datum.GetFloat("X")
	var y = datum.GetFloat("Y")
	var angle = datum.GetFloat("Angle")
	p.SetPosition(types.NewVector(x, y))
	p.SetAngle(angle)
}

const Class types.Channel = 0x82
