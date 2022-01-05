package wall

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
)

type Wall struct {
	entities.ID
	collider.Collider

	width int
}

func NewWall(w int) *Wall {
	var wall = &Wall{
		width: w,
	}
	wall.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(0, 0), 0, float64(w), 10),
		},
		[]collider.Circle{},
		10,
	)
	return wall
}
