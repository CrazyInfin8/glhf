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

func NewFrameFromImage(path AssetPath, cache, unique bool) (*Frame, error) {
	graphic, err := assets.LoadImage(path, true, false)
	if err != nil {
		return nil, err
	}

	f := new(Frame)
	f.graphic = graphic
	return f, nil
}

func (f *Frame) Width() int {
	return f.graphic.Bounds().Dx()
}

func (f *Frame) Height() int {
	return f.graphic.Bounds().Dy()
}
