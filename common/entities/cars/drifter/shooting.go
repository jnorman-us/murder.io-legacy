package drifter

const NumberBullets = 5
const BulletSpeed = 10
const DamagePerBullet = 5
const DropOff = 50

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

func (d *Drifter) GetDamagePerBullet() int {
	return DamagePerBullet
}

func (d *Drifter) GetBulletDropOff() float64 {
	return DropOff
}

func (d *Drifter) GetBulletSpeed() float64 {
	return BulletSpeed
}
