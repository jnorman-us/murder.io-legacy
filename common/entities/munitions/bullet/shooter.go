package bullet

import "github.com/josephnormandev/murder/common/types"

type Shooter interface {
	GetID() types.ID
	GetPosition() types.Vector
	GetDamagePerBullet() int
	GetBulletDropOff() float64
	GetBulletSpeed() float64
}

func (b *Bullet) GetShooter() types.ID {
	var shooter = *b.shooter
	return shooter.GetID()
}
