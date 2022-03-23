package pole

import (
	"github.com/josephnormandev/murder/common/packets"
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

const Class types.Channel = 0x83
