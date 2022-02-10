package drawer

import (
	"github.com/josephnormandev/murder/common/types"
	draw "github.com/llgcode/draw2d/draw2dimg"
)

type Drawable interface {
	// Draw(*draw.GraphicContext)
	DrawHitbox(*draw.GraphicContext)
}

func (d *Drawer) AddDrawable(id types.ID, w *Drawable) {
	d.Drawables[id] = w
}

func (d *Drawer) RemoveDrawable(id types.ID) {
	delete(d.Drawables, id)
}
