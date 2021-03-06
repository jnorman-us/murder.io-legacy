package types

import "image/color"

type colors struct {
	Red    color.RGBA
	Green  color.RGBA
	Yellow color.RGBA
	Orange color.RGBA
	Blue   color.RGBA
	Gray   color.RGBA
}

var Colors = colors{
	Red:    color.RGBA{R: 0xff, A: 0xff},
	Green:  color.RGBA{G: 0xff, A: 0xff},
	Yellow: color.RGBA{R: 0xff, G: 0xff, A: 0xff},
	Orange: color.RGBA{R: 0xff, G: 0x88, A: 0xff},
	Blue:   color.RGBA{B: 0xff, A: 0xff},
	Gray:   color.RGBA{R: 0x60, G: 0x60, B: 0x60, A: 0xff},
}
