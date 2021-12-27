package sword

import "github.com/josephnormandev/murder/common/types"

type Beholder interface {
	GetID() int
	GetPosition() types.Vector
	GetVelocity() types.Vector
}
