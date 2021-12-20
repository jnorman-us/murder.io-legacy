package engine

type Engine struct {
	running   bool
	Moveables map[int]*Moveable
}

func NewEngine() *Engine {
	var engine = &Engine{
		running:   false,
		Moveables: map[int]*Moveable{},
	}
	return engine
}

func (e *Engine) UpdatePhysics() {
	for id := range e.Moveables {
		(*e.Moveables[id]).UpdatePosition()
	}
}
