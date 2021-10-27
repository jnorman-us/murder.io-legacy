package engine

import (
	"github.com/josephnormandev/murder/common/collisions"
	"github.com/josephnormandev/murder/common/world"
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
	go e.run()
}

func (e *Engine) run() {
	// loop for each tick of the game
	for range time.Tick(time.Second / 20) { // 20 TPS
		if e.running == false {
			break
		}
		e.Tick()
	}
}

func (e *Engine) Tick() {
	for _, moveable := range e.world.Moveables {
		(*moveable).UpdatePosition()
	}

	e.collisionsManager.Resolve()

	for _, identifiable := range e.world.Identifiables {
		(*identifiable).Tick()
	}
}

func (e *Engine) Stop() {
	e.running = false
}
