package glhf

import "image/color"

type Color uint32

const (
	Red = 0xFFFF0000
	Green = 0xFF00FF00
)

func (c Color) RGBA() color.RGBA {
	return color.RGBA{
		A: uint8(c >> 24),
		R: uint8(c >> 16),
		G: uint8(c >> 8),
		B: uint8(c),
	}
}