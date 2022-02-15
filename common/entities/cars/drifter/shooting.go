package drifter

import "github.com/josephnormandev/murder/common/types"

const NumberBullets = 10
const BulletSpeed = 15
const DamagePerBullet = 9
const DropOff = 300

func (d *Drifter) Shoot() {
	const fanAngle = .02
	var angle = d.GetAngle()
	angle -= (NumberBullets * fanAngle) / 2

	var spawner = *d.spawner
	for i := 0; i < NumberBullets; i++ {
		spawner.DrifterShootBullet(d, angle)
		angle += fanAngle
	}
}

func (d *Drifter) GetShootingCoolDown() *types.CoolDown {
	return &d.shotgunCoolDown
}

func (d *Drifter) GetDamagePerBullet() int {
	return DamagePerBullet
}

func (d *Drifter) GetBulletDropOff() float64 {
	return DropOff
}

func (d *Drifter) GetBulletSpeed() float64 {
	return BulletSpeed
}
