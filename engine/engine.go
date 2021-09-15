package engine

import (
	"github.com/josephnormandev/murder/collisions"
	"github.com/josephnormandev/murder/types"
	"github.com/josephnormandev/murder/world"
	"time"
)

type Engine struct {
	running           bool
	world             *world.World
	collisionsManager *collisions.Manager
}

func NewEngine(w *world.World) *Engine {
	var engine = &Engine{
		running:           false,
		world:             w,
		collisionsManager: w.CollisionsManager,
	}
	return engine
}

func (e *Engine) Start() {
	e.running = true
	e.run()
}

func (e *Engine) run() {
	// loop for each tick of the game
	for range time.Tick(time.Second / 20) { // 20 TPS
		if e.running == false {
			break
		}
		// update position based on Force and Velocity
		for _, moveable := range e.world.Moveables {
			(*moveable).UpdatePosition()
		}

		for a, actorA := range e.collisionsManager.ActorActorCollidables {
			var colliderA = (*actorA).GetCollider()
			colliderA.SetBGColor(types.Default)
			for b, actorB := range e.collisionsManager.ActorActorCollidables {
				if a != b {
					var colliderB = (*actorB).GetCollider()

					if colliderA.CheckCollision(colliderB) {
						(*actorA).ActorCollidedWithActor(actorB)
					}
				}
			}
		}

		for _, wall := range e.collisionsManager.WallActorCollidables {
			var colliderA = (*wall).GetCollider()
			colliderA.SetBGColor(types.Default)
			for _, actor := range e.collisionsManager.ActorWallCollidables {
				var colliderB = (*actor).GetCollider()

				if colliderA.CheckCollision(colliderB) {
					(*actor).ActorCollidedWithWall(wall)
					(*wall).WallCollidedWithActor(actor)
				}
			}
		}

		for _, identifiable := range e.world.Identifiables {
			(*identifiable).Tick()
		}
	}
}

func (e *Engine) Stop() {
	e.running = false
}
