package dimetrodon

import (
	"github.com/josephnormandev/murder/common/packets"
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

const Class types.Channel = 0x82
