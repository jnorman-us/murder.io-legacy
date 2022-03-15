package dimetrodon

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/communications/data/schemas"
	"github.com/josephnormandev/murder/common/types"
)

func (d *Dimetrodon) GetClass() types.Channel {
	return Class
}

func (d *Dimetrodon) GetData() data.Data {
	var datum = data.NewData(schemas.DimetrodonSchema)
	datum = d.ColliderData(datum)
	datum.SetInteger("Health", d.Health.Health)
	datum.SetInteger("MaxHealth", d.Health.MaxHealth)
	datum.SetString("Username", string(d.UserID))
	return datum
}

func (d *Dimetrodon) FromData(datum data.Data) {
	datum.ApplySchema(schemas.DimetrodonSchema)
	d.ColliderFromData(datum)
}

const Class types.Channel = 0x84
