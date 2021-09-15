package entities

import (
	"fmt"
	"github.com/josephnormandev/murder/classes"
	"github.com/josephnormandev/murder/collider"
	"github.com/josephnormandev/murder/collisions"
	"github.com/josephnormandev/murder/types"
	"github.com/josephnormandev/murder/world"
)

type Walls struct {
	Entity
	layout [][]bool
}

func NewWalls(l [][]bool) *Walls {
	const squareWidth = 1
	const squareHeight = 1

	var walls = &Walls{
		layout: l,
	}
	var rectangles []collider.Rectangle
	for y := range l {
		for x := range l[y] {
			if l[y][x] {
				var xPos = float64(x * squareWidth)
				var yPos = float64(y * squareHeight)

				rectangles = append(
					rectangles,
					collider.NewRectangle(
						types.NewVector(xPos, yPos),
						0, //math.Pi / 2,
						squareWidth,
						squareHeight,
					),
				)
			}
		}
	}

	walls.Setup(
		rectangles,
		[]collider.Circle{},
	)
	fmt.Println(rectangles)
	return walls
}

func (w *Walls) AddTo(wo *world.World) {
	var id = wo.NextAvailableID()
	w.SetID(id)
	w.world = wo

	var drawable = classes.Drawable(w)
	var moveable = classes.Moveable(w)
	var wallActor = collisions.WallActorCollidable(w)

	wo.AddDrawable(id, &drawable)
	wo.AddMoveable(id, &moveable)
	wo.CollisionsManager.AddWallActor(id, &wallActor)
}

func (w *Walls) RemoveFrom() {
	var wo = w.world
	var id = w.GetID()

	wo.RemoveDrawable(id)
	wo.RemoveMoveable(id)
	wo.CollisionsManager.RemoveWallActor(id)
}

func (w *Walls) Tick() {

}

func (w *Walls) WallCollidedWithActor(a *collisions.ActorWallCollidable) {
	w.SetBGColor(types.White)
}
