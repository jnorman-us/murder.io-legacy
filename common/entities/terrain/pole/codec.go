package pole

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/communications/data/schemas"
	"github.com/josephnormandev/murder/common/types"
)

func (p *Pole) GetClass() types.Channel {
	return 0x82
}

func (p *Pole) GetData() data.Data {
	var datum = data.NewData(schemas.PoleSchema)
	datum = p.ColliderData(datum)
	return datum
}

func (p *Pole) FromData(datum data.Data) {
	datum.ApplySchema(schemas.PoleSchema)
	p.ColliderFromData(datum)
}

const Class types.Channel = 0x82
