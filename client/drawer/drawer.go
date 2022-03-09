package drawer

import (
	"github.com/josephnormandev/murder/common/types"
	"syscall/js"
	"time"
)

type Drawer struct {
	Drawables  map[types.ID]*Drawable
	Centerable *Centerable

	lastStart       time.Time
	lastDuration    float64
	averageDuration float64
	update          func(float64)
}

func NewDrawer() *Drawer {
	var drawer = &Drawer{
		Drawables: map[types.ID]*Drawable{},
	}

	return drawer
}

func (d *Drawer) GetDrawData(this js.Value, values []js.Value) interface{} {
	return nil
}
