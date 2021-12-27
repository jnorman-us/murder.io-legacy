package drawer

import "github.com/josephnormandev/murder/common/types"

type Centerable interface {
	GetPosition() types.Vector
}

func (d *Drawer) SetCenterable(c *Centerable) {
	d.center = c
}
