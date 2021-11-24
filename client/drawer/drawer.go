package drawer

import (
	"github.com/josephnormandev/murder/common/engine"
	"github.com/josephnormandev/murder/common/world"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
	"image/color"
)

type Drawer struct {
	world  *world.World
	engine *engine.Engine
	canvas *canvas.Canvas2d
}

func NewDrawer(w *world.World, e *engine.Engine, width, height int) *Drawer {
	var c, _ = canvas.NewCanvas2d(false)
	c.Create(
		width,
		height,
	)

	return &Drawer{
		world:  w,
		engine: e,
		canvas: c,
	}
}

func (d *Drawer) Start() {
	d.canvas.Start(120, d.Render) // random maxFPS, change to some setting later?
}

func (d *Drawer) Render(g *draw2dimg.GraphicContext) bool {
	d.engine.Tick()

	g.SetFillColor(color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})
	g.Clear()

	for _, collidable := range d.world.CollisionsManager.Collidables {
		(*collidable).GetCollider().Draw(g)
	}

	return true
}
