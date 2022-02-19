package bullet

import (
	"github.com/josephnormandev/murder/common/collider"
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
		map[string]collider.Rectangle{
			"body": collider.NewRectangle(types.NewVector(0, 0), 1, 10, 10),
		},
		map[string]collider.Circle{
			//"body": collider.NewCircle(types.NewVector(0, 0), 5),
		},
	)
	b.Collider.SetColor(types.Colors.Blue)
	b.Collider.SetMass(Mass)
	b.Collider.SetForwardFriction(Friction)
	if b.shooter != nil {
		var shooter = *b.shooter
		b.SetAngle(b.Angle + shooter.GetAngle())
		b.initialPosition = shooter.GetPosition()
		b.damage = shooter.GetDamagePerBullet()
		b.dropOff = shooter.GetBulletDropOff()

		var bulletVelocity = types.NewVector(shooter.GetBulletSpeed(), 0)
		bulletVelocity.RotateAbout(b.GetAngle(), types.NewZeroVector())
		var velocity = shooter.GetVelocity()
		velocity.Add(bulletVelocity)
		b.Collider.SetPosition(b.initialPosition)
		b.Collider.SetVelocity(bulletVelocity)
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

func (b *Bullet) Break() {
	var spawner = *b.spawner
	spawner.RemoveBullet(b.ID)
}

func (b *Bullet) Stop() {
	var spawner = *b.spawner
	spawner.DisableBullet(b.ID)
}
