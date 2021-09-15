package drawer

import (
	"github.com/josephnormandev/murder/types"
	world2 "github.com/josephnormandev/murder/world"
	"github.com/nsf/termbox-go"
	"math"
	"time"
)

type Drawer struct {
	running bool
	world   *world2.World
}

func NewDrawer(w *world2.World) *Drawer {
	return &Drawer{
		running: false,
		world:   w,
	}
}

func (d *Drawer) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	d.running = true
	d.run()
}

func (d *Drawer) run() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for range time.Tick(time.Second / 20) { // 20 FPS
		if d.running == false {
			break
		}

		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		for _, drawable := range d.world.Drawables {
			(*drawable).Draw(setCell)
		}

		if err := termbox.Flush(); err != nil {
			panic(err)
		}
	}
}

func (d *Drawer) Stop() {
	d.running = false
}

func setCell(v types.Vector, c rune, background types.Color, text types.Color) {
	var backgroundColor = termbox.Attribute(background)
	var textColor = termbox.Attribute(text)

	termbox.SetCell(int(math.Floor(v.X)*3), int(math.Floor(v.Y)), c, textColor, backgroundColor)
	termbox.SetCell(int(math.Floor(v.X)*3+1), int(math.Floor(v.Y)), c, textColor, backgroundColor)
	termbox.SetCell(int(math.Floor(v.X)*3+2), int(math.Floor(v.Y)), c, textColor, backgroundColor)
}
