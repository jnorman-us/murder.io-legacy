package innocent

// Swingable is an abstraction of a sword which is wielded by the
// player and swung around
type Swingable interface {
	GetID() int
	Swing()
	SwingCompleted() bool
}
