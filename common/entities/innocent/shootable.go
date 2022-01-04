package innocent

type Shootable interface {
	GetID() int
	Charge()
	Fire()
	Fired() bool
	Cancel()
}
