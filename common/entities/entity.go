package entities

import (
	"github.com/josephnormandev/murder/common/collider"
	"github.com/josephnormandev/murder/common/world"
)

type Entity struct {
	collider.Collider
	id     int32
	health float64
	maxHealth float64 // minHealth is obviously 0
	maxSpeed float64
	world  *world.World
}

func (e *Entity) SetStats(maxHealth, maxSpeed float64) {
	e.maxHealth = maxHealth
	e.maxSpeed = maxSpeed
}

func (e *Entity) GetID() int32 {
	return e.id
}

func (e *Entity) SetID(id int32) {
	e.id = id
}

func (e *Entity) GetMaxHealth() float64 {
	return e.maxHealth
}

func (e *Entity) GetMaxSpeed() float64 {
	return e.maxSpeed
}

func (e *Entity) GetHealth() float64 {
	return e.health
}

func (e *Entity) SetHealth(newHealth float64) {
	e.health = newHealth
}