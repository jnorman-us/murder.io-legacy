package engine

type Moveable interface {
	GetID() int
	Tick()
	UpdatePosition()
}

func (e *Engine) AddMoveable(id int, m *Moveable) {
	e.Moveables[id] = m
}

func (e *Engine) RemoveMoveable(id int) {
	delete(e.Moveables, id)
}
