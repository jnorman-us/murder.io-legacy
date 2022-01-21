package wall

import (
	"encoding/gob"
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/entities"
	"github.com/josephnormandev/murder/common/types"
)

type Wall struct {
	entities.ID
	collider.Collider

	Width int
}

func NewWall(w int) *Wall {
	var wall = &Wall{
		Width: w,
	}
	wall.Setup()
	return wall
}

func (w *Wall) Setup() {
	w.SetupCollider(
		[]collider.Rectangle{
			collider.NewRectangle(types.NewVector(0, 0), 0, float64(w.Width), 10),
		},
		[]collider.Circle{},
		10,
	)
	w.SetColor(types.Colors.Gray)
}

func (w *Wall) GetClass() string {
	return "wall"
}

func (w *Wall) GetData(e *gob.Encoder) error {
	return e.Encode(w)
}
