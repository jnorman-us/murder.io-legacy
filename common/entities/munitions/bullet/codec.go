package bullet

import (
	"github.com/josephnormandev/murder/common/packets"
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

const Class types.Channel = 0x81
