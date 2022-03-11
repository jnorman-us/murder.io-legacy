package dimetrodon

import "github.com/josephnormandev/murder/common/types"

const BulletSpeed = 5
const DamagePerBullet = 9
const DropOff = 500

func (d *Dimetrodon) Shoot() {
	var spawner = *d.spawner
	spawner.DimetrodonShootBullet(d)
}

func (d *Dimetrodon) GetShootingCoolDown() *types.CoolDown {
	return &d.gattlingGunCoolDown
}

func (d *Dimetrodon) GetDamagePerBullet() int {
	return DamagePerBullet
}

func (d *Dimetrodon) GetBulletDropOff() float64 {
	return DropOff
}

func (d *Dimetrodon) GetBulletSpeed() float64 {
	return BulletSpeed
}
