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
	Damage  int
	DropOff float64
	shooter *Shooter
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
		b.Damage = shooter.GetDamagePerBullet()
		b.DropOff = shooter.GetBulletDropOff()
		b.Position = shooter.GetPosition()

		var velocity = types.NewVector(shooter.GetBulletSpeed(), 0)
		velocity.RotateAbout(b.GetAngle(), types.NewZeroVector())
		b.Collider.SetVelocity(velocity)
	}
}
