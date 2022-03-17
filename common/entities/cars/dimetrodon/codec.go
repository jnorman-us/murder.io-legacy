package dimetrodon

import (
	"github.com/josephnormandev/murder/common/communications/data"
	"github.com/josephnormandev/murder/common/communications/data/schemas"
	"github.com/josephnormandev/murder/common/types"
)

func (d *Dimetrodon) GetClass() types.Channel {
	return Class
}

func (d *Dimetrodon) GetStartData() data.Data {
	var datum = data.NewData(schemas.DimetrodonStartSchema)
	datum.SetFloat("X", d.Position.X)
	datum.SetFloat("Y", d.Position.Y)
	datum.SetFloat("Angle", d.Angle)
	datum.SetInteger("Health", d.Health.Health)
	datum.SetString("Username", string(d.UserID))
	return datum
}

func (d *Dimetrodon) GetData() data.Data {
	var datum = data.NewData(schemas.DimetrodonSchema)
	datum.SetInteger("Health", d.Health.Health)
	return datum
}

func (d *Dimetrodon) FromData(datum data.Data) {
	datum.ApplySchema(schemas.DimetrodonSchema)
	d.SetHealth(datum.GetInteger("Health"))
}

func (d *Dimetrodon) FromStartData(datum data.Data) {
	datum.ApplySchema(schemas.DimetrodonStartSchema)
	var x = datum.GetFloat("X")
	var y = datum.GetFloat("Y")
	var angle = datum.GetFloat("Angle")
	var health = datum.GetInteger("Health")
	var username = datum.GetString("Username")
	d.SetPosition(types.NewVector(x, y))
	d.SetAngle(angle)
	d.SetHealth(health)
	d.UserID = types.UserID(username)
}

const Class types.Channel = 0x84
