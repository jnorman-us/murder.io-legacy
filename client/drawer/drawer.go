package drawer

import (
	"encoding/json"
	"fmt"
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

func (d *Drawer) DrawUpdate(this js.Value, values []js.Value) interface{} {
	var objects = map[types.ID]DrawableObject{}

	for id, d := range d.Drawables {
		var drawable = *d
		objects[id] = DrawableObject{
			ID:       id,
			Position: drawable.GetPosition(),
			Angle:    drawable.GetAngle(),
			Color:    drawable.GetColor(),
		}
	}

	bytes, err := json.Marshal(objects)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return string(bytes)
}

func (d *Drawer) CenterUpdate(this js.Value, values []js.Value) interface{} {
	var centerObj = CenterObject{}
	if d.Centerable != nil {
		var centerable = *d.Centerable
		centerObj.Position = centerable.GetPosition()
		centerObj.Angle = centerable.GetAngle()
	}

	bytes, err := json.Marshal(centerObj)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return string(bytes)
}
