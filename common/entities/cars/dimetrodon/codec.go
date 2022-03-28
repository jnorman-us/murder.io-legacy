package dimetrodon

import (
	"github.com/josephnormandev/murder/common/packets"
	"github.com/josephnormandev/murder/common/packets/schemas"
	"github.com/josephnormandev/murder/common/types"
)

func (d *Dimetrodon) GetClass() types.Channel {
	return Class
}

func (d *Dimetrodon) PopulateData(data *packets.Data) {
	data.SetFloat("X", d.Position.X)
	data.SetFloat("Y", d.Position.Y)
	data.SetFloat("Angle", d.Angle)
	data.SetInteger("Health", d.Health.Health)
	data.SetInteger("MaxHealth", d.MaxHealth)
	data.SetString("Username", string(d.UserID))
}

func (d *Dimetrodon) FromData(data packets.Data) {
	var position = types.NewVector(data.GetFloat("X"), data.GetFloat("Y"))
	var angle = data.GetFloat("Angle")
	d.SetPosition(position)
	d.SetAngle(angle)

	d.SetHealth(data.GetInteger("Health"))
	d.Health.MaxHealth = data.GetInteger("MaxHealth")
	d.UserID = types.UserID(data.GetString("Username"))
}

var Class = schemas.DimetrodonSchema.Channel()
