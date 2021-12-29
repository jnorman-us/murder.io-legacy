package drawer

import (
	"github.com/josephnormandev/murder/common/types"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
	"image/color"
	"syscall/js"
	"time"
)

type Drawer struct {
	canvas           *canvas.Canvas2d
	canvasDimensions types.Vector

	center    *Centerable
	Drawables map[int]*Drawable

	lastStart     time.Time
	lastDuration  float64
	updatePhysics func(float64)
}

func NewDrawer() *Drawer {
	var window = js.Global()

	var c, _ = canvas.NewCanvas2d(false)
	c.Create(
		window.Get("innerWidth").Int(),
		window.Get("innerHeight").Int(),
	)

	return &Drawer{
		canvas:           c,
		canvasDimensions: types.NewVector(float64(c.Width()), float64(c.Height())),
		Drawables:        map[int]*Drawable{},
	}
}

func (d *Drawer) Start(updatePhysics func(float64)) {
	d.updatePhysics = updatePhysics
	d.lastStart = time.Now()
	d.lastDuration = 10

	d.canvas.Start(200, d.render) // random maxFPS, change to some setting later?
}

func (d *Drawer) render(g *draw2dimg.GraphicContext) bool {
	d.lastDuration = float64(time.Since(d.lastStart).Milliseconds())
	d.lastStart = time.Now()
	d.updatePhysics(d.lastDuration)

	// fmt.Println("FPS:", 1000 / d.lastDuration)

	var translated types.Vector
	if d.center != nil {
		var position = (*d.center).GetPosition()
		position.Scale(-1) // negate to subtract

		translated = d.canvasDimensions
		translated.Scale(.5)
		translated.Add(position)

		g.Translate(translated.X, translated.Y)
	}
	g.SetFillColor(color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})
	g.Clear()

	for _, drawable := range d.Drawables {
		(*drawable).DrawHitbox(g)
	}

	if d.center != nil {
		translated.Scale(-1)
		g.Translate(translated.X, translated.Y)
	}

	return true
}

func (d *Drawer) GetDimensions() types.Vector {
	return d.canvasDimensions
}
