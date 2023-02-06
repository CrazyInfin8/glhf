package glhf

import (
	"image/color"

	"github.com/crazyinfin8/glhf/driver"
)

type (
	Frame struct {
		graphic      Graphic
		flipX, flipY bool
	}
)

func NewFrameWithColor(width, height int, c color.Color) *Frame {
	f := new(Frame)
	f.graphic = driver.Drivers.NewGraphic(width, height)
	f.graphic.Fill(c)

	return f
}

func (f *Frame) Width() int {
	return f.graphic.Bounds().Dx()
}

func (f *Frame) Height() int {
	return f.graphic.Bounds().Dy()
}