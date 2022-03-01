package drawer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/josephnormandev/murder/common/types"
	"time"
)

type Drawer struct {
	Drawables  map[types.ID]*Drawable
	images     map[types.ID]*ebiten.Image
	Centerable *Centerable

	lastStart       time.Time
	lastDuration    float64
	averageDuration float64
	update          func(float64)
}

func NewDrawer() *Drawer {
	var drawer = &Drawer{
		Drawables: map[types.ID]*Drawable{},
		images:    map[types.ID]*ebiten.Image{},
	}

	return drawer
}

func (d *Drawer) Start(update func(float64)) {
	d.update = update
	d.lastStart = time.Now()
	d.lastDuration = 1000 / 60 // duration of 60fps frame
	d.averageDuration = d.lastDuration

	var eGame = ebiten.Game(d)
	ebiten.SetRunnableOnUnfocused(true)
	err := ebiten.RunGame(eGame)
	if err != nil {
		return
	}
}

func (d *Drawer) Update() error {
	d.lastDuration = float64(time.Since(d.lastStart).Milliseconds())
	d.averageDuration = (1*d.averageDuration + .02*d.lastDuration) / 1.02
	d.lastStart = time.Now()
	d.update(d.averageDuration)
	return nil
}

func (d *Drawer) Draw(screen *ebiten.Image) {
	var screenX, screenY = screen.Size()

	var center = types.NewZeroVector()
	var centerTranslate = types.NewZeroVector()
	var centerRotate = 0.0
	if d.Centerable != nil {
		var centerable = *d.Centerable

		var centerOfScreen = types.NewVector(float64(screenX), float64(screenY))
		centerOfScreen.Scale(.5)

		center = centerable.GetPosition()
		centerRotate = -centerable.GetAngle()

		centerTranslate = centerOfScreen
		var centerCopy = center
		centerCopy.Scale(-1)
		centerTranslate.Add(centerCopy)
	}

	for id, _ := range d.Drawables {
		var drawable = *d.Drawables[id]
		var image = d.images[id]

		var angle = drawable.GetAngle()
		var position = drawable.GetPosition()

		var centeredPosition = position
		centeredPosition.RotateAbout(centerRotate, center)
		centeredPosition.Add(centerTranslate)

		var options = &ebiten.DrawImageOptions{}
		options.GeoM.Rotate(angle + centerRotate)
		options.GeoM.Translate(centeredPosition.X, centeredPosition.Y)

		screen.DrawImage(image, options)
	}
}

func (d *Drawer) Layout(outW, outH int) (int, int) {
	return 1200, 800
}
