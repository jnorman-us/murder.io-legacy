package collider

import "github.com/josephnormandev/murder/common/types"

type Collision struct {
	colliding     bool
	contactPoints []types.Vector
}

func NewCollision() Collision {
	return Collision{
		colliding:     false,
		contactPoints: []types.Vector{},
	}
}

func (c *Collision) Colliding() bool {
	return c.colliding
}

func (c *Collision) SetColliding(colliding bool) {
	c.colliding = true
}

func (c *Collision) AddContactPoint(point types.Vector) {
	c.contactPoints = append(c.contactPoints, point)
}

func (c *Collision) GetContactPoint() types.Vector {
	var average = types.NewZeroVector()

	for _, point := range c.contactPoints {
		average.Add(point)
	}

	average.Scale(float64(len(c.contactPoints)))
	return average
}
