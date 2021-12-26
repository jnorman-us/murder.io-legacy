package drawer

import (
	draw "github.com/llgcode/draw2d/draw2dimg"
)

type Drawable interface {
	// Draw(*draw.GraphicContext)
	DrawHitbox(*draw.GraphicContext)
}

func (d *Drawer) AddDrawable(id int, w *Drawable) {
	d.Drawables[id] = w
}

func (d *Drawer) RemoveDrawable(id int) {
	delete(d.Drawables, id)
}
