package types

import (
	"github.com/nsf/termbox-go"
	"math/rand"
)

type Color termbox.Attribute

const (
	Default = Color(termbox.ColorDefault)
	Red     = Color(termbox.ColorRed)
	Blue    = Color(termbox.ColorBlue)
	Green   = Color(termbox.ColorGreen)
	Black   = Color(termbox.ColorBlack)
	White   = Color(termbox.ColorWhite)
)

func RandomColor() Color {
	var colors = []termbox.Attribute{
		termbox.ColorBlue,
		termbox.ColorRed,
		termbox.ColorGreen,
		termbox.ColorCyan,
		termbox.ColorMagenta,
		termbox.ColorWhite,
		termbox.ColorYellow,
	}
	return Color(colors[rand.Intn(len(colors))])
}
