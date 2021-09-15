package classes

import "github.com/josephnormandev/murder/types"

type Drawable interface {
	Draw(func(types.Vector, rune, types.Color, types.Color))
}
