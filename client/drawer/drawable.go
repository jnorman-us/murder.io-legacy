package drawer

import (
	"github.com/josephnormandev/murder/common/types"
	"image/color"
)

type Drawable interface {
	GetID() types.ID
	GetPosition() types.Vector
	GetAngle() float64
	GetColor() color.RGBA
}

func (d *Drawer) AddDrawable(id types.ID, w *Drawable) {
	d.Drawables[id] = w
}

func (d *Drawer) RemoveDrawable(id types.ID) {
	delete(d.Drawables, id)
}
