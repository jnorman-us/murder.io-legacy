package innocent

type Shootable interface {
	GetID() int
	ChargeBow()
	Fire()
	IsFired() bool
	Cancel()
}
