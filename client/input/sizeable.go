package input

import "github.com/josephnormandev/murder/common/types"

type Sizeable interface {
	GetDimensions() types.Vector
}
