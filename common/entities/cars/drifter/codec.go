package drifter

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/communications/data/schemas"
	"github.com/josephnormandev/murder/common/types"
)

func (d *Drifter) GetClass() types.Channel {
	return Class
}

func (d *Drifter) GetData() data.Data {
	var datum = data.NewData(schemas.DrifterSchema)
	datum = d.ColliderData(datum)
	datum.SetInteger("Health", d.Health.Health)
	datum.SetInteger("MaxHealth", d.Health.MaxHealth)
	datum.SetString("Username", string(d.UserID))
	return datum
}

func (d *Drifter) FromData(datum data.Data) {
	datum.ApplySchema(schemas.DrifterSchema)
	d.ColliderFromData(datum)
}

const Class types.Channel = 0x81
