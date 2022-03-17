package world

import (
	"github.com/josephnormandev/murder/common/entities/cars/dimetrodon"
	"github.com/josephnormandev/murder/common/entities/munitions/bullet"
)

func (w *World) DimetrodonShootBullet(d *dimetrodon.Dimetrodon) {
	var shooter = bullet.Shooter(d)
	var newBullet = bullet.NewBullet(&shooter, 0)
	w.AddBullet(newBullet)
}
