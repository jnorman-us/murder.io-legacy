package drawer

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
	"image/color"
)

type Drawer struct {
	canvas    *canvas.Canvas2d
	Drawables map[int]*Drawable
}

func NewDrawer(width, height int) *Drawer {
	var c, _ = canvas.NewCanvas2d(false)
	c.Create(
		width,
		height,
	)

	return &Drawer{
		canvas:    c,
		Drawables: map[int]*Drawable{},
	}
}

func (d *Drawer) Start() {
	d.canvas.Start(120, d.Render) // random maxFPS, change to some setting later?
}

func (d *Drawer) Render(g *draw2dimg.GraphicContext) bool {
	g.SetFillColor(color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})
	g.Clear()

	for _, drawable := range d.Drawables {
		(*drawable).DrawHitbox(g)
	}

	return true
}
