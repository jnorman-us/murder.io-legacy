package drawer

import (
	"github.com/josephnormandev/murder/common/types"
	"image/color"
)

type DrawableObject struct {
	ID       types.ID
	Position types.Vector
	Angle    float64
	Color    color.RGBA
}

type CenterObject struct {
	Position types.Vector
	Angle    float64
}
