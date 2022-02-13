package bullet

import (
	"github.com/josephnormandev/murder/common/collisions/collider"
	"github.com/josephnormandev/murder/common/types"
)

const Mass = 1
const Friction = 0

type Bullet struct {
	types.ID
	collider.Collider

	initialPosition types.Vector
	damage          int
	dropOff         float64

	shooter *Shooter
	spawner *Spawner
}

func NewBullet(s *Shooter, angle float64) *Bullet {
	var bullet = &Bullet{
		shooter: s,
	}
	bullet.SetAngle(angle)
	bullet.Setup()
	return bullet
}

func (b *Bullet) Setup() {
	b.SetupCollider(
		[]collider.Rectangle{},
		[]collider.Circle{
			collider.NewCircle(types.NewVector(0, 0), 5),
		},
	)
	b.Collider.SetColor(types.Colors.Blue)
	b.Collider.SetMass(Mass)
	b.Collider.SetFriction(Friction)
	if b.shooter != nil {
		var shooter = *b.shooter
		b.initialPosition = shooter.GetPosition()
		b.damage = shooter.GetDamagePerBullet()
		b.dropOff = shooter.GetBulletDropOff()

		var velocity = types.NewVector(shooter.GetBulletSpeed(), 0)
		velocity.RotateAbout(b.GetAngle(), types.NewZeroVector())
		b.Collider.SetPosition(b.initialPosition)
		b.Collider.SetVelocity(velocity)
	}
}

func (b *Bullet) GetRange() float64 {
	return b.dropOff
}

func (b *Bullet) GetInitialPosition() types.Vector {
	return b.initialPosition
}

func (b *Bullet) GetDamage() int {
	return b.damage
}

func (b *Bullet) Destroy() {
	var spawner = *b.spawner
	spawner.RemoveBullet(b.ID)
}
