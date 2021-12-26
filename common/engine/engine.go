package engine

type Engine struct {
	Moveables map[int]*Moveable
}

func NewEngine() *Engine {
	var engine = &Engine{
		Moveables: map[int]*Moveable{},
	}
	return engine
}

func (e *Engine) UpdatePhysics() {
	for id := range e.Moveables {
		(*e.Moveables[id]).UpdatePosition()
	}
}
