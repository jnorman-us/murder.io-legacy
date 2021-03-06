package drawer

import "github.com/josephnormandev/murder/common/types"

type Centerable interface {
	GetPosition() types.Vector
	GetAngle() float64
}

func (d *Drawer) SetCenterable(c *Centerable) {
	d.Centerable = c
}
