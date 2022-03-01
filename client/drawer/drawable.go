package drawer

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	var color = (*w).GetColor()
	var image = ebiten.NewImage(25, 25)
	image.Fill(color)

	d.Drawables[id] = w
	d.images[id] = image
}

func (d *Drawer) RemoveDrawable(id types.ID) {
	delete(d.Drawables, id)
	delete(d.images, id)
}
