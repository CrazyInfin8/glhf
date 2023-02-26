package color

import "image/color"

type Color64 uint64

func (c Color64) RGBA() (r, g, b, a uint32) {
	return color.RGBA64{
		R: uint16(c >> 0),
		G: uint16(c >> 0),
		B: uint16(c >> 0),
		A: uint16(c >> 0),
	}.RGBA()
}
