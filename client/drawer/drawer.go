package drawer

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/josephnormandev/murder/common/types"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
	"log"
	"syscall/js"
	"time"
)

type Drawer struct {
	canvas           *canvas.Canvas2d
	canvasDimensions types.Vector

	center    *Centerable
	Drawables map[types.ID]*Drawable

	lastStart       time.Time
	lastDuration    float64
	averageDuration float64
	updatePhysics   func(float64)

	fontdata draw2d.FontData
}

func NewDrawer() *Drawer {
	var window = js.Global()

	var c, _ = canvas.NewCanvas2d(false)
	c.Create(
		window.Get("innerWidth").Int(),
		window.Get("innerHeight").Int(),
	)

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	var drawer = &Drawer{
		canvas:           c,
		canvasDimensions: types.NewVector(float64(c.Width()), float64(c.Height())),
		Drawables:        map[types.ID]*Drawable{},
		fontdata:         draw2d.FontData{Name: "goregular", Family: draw2d.FontFamilyMono, Style: draw2d.FontStyleNormal},
	}

	draw2d.RegisterFont(
		drawer.fontdata,
		font,
	)

	return drawer
}

func (d *Drawer) Start(updatePhysics func(float64)) {
	d.updatePhysics = updatePhysics
	d.lastStart = time.Now()
	d.lastDuration = 1000 / 60 // duration of 60fps frame
	d.averageDuration = d.lastDuration

	d.canvas.Start(200, d.render) // random maxFPS, change to some setting later?
}

func (d *Drawer) GetFPS() int {
	return int(1000 / d.averageDuration)
}

func (d *Drawer) render(g *draw2dimg.GraphicContext) bool {
	d.lastDuration = float64(time.Since(d.lastStart).Milliseconds())
	d.averageDuration = (1*d.averageDuration + .02*d.lastDuration) / 1.02
	d.lastStart = time.Now()
	d.updatePhysics(d.averageDuration)

	g.SetFillColor(color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})
	g.Clear()
	d.drawFPS(g)

	var translated types.Vector
	if d.center != nil {
		var position = (*d.center).GetPosition()
		position.Scale(-1) // negate to subtract

		translated = d.canvasDimensions
		translated.Scale(.5)
		translated.Add(position)

		g.Translate(translated.X, translated.Y)
	}

	for _, drawable := range d.Drawables {
		(*drawable).DrawHitbox(g)
	}

	if d.center != nil {
		translated.Scale(-1)
		g.Translate(translated.X, translated.Y)
	}

	return true
}

func (d *Drawer) drawFPS(gc *draw2dimg.GraphicContext) {
	gc.SetFontData(d.fontdata)

	// Set the fill text color to black
	gc.SetFillColor(image.Black)
	gc.SetStrokeColor(image.Black)
	gc.SetFontSize(10)

	// Draw Text
	gc.FillStringAt(fmt.Sprintf("FPS: %d", d.GetFPS()), 8, 20)
	gc.FillStringAt("WASD - Movement", 8, 34)
	gc.FillStringAt("Left Click - Swing Sword", 8, 48)
	gc.FillStringAt("Right Click - Shoot Arrow (hold to charge)", 8, 62)
}

func (d *Drawer) GetDimensions() types.Vector {
	return d.canvasDimensions
}
