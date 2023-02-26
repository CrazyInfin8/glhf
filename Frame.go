package glhf

import (
	"image/color"

	"github.com/crazyinfin8/glhf/driver"
)

type (
	Frame struct {
		_graphic
		flipX, flipY bool
	}
)

func NewFrameWithColor(width, height int, c color.Color) *Frame {
	f := new(Frame)

	graphic := driver.Drivers.NewGraphic(width, height, driver.DefaultGraphicOptions())
	graphic.Fill(c)
	f._graphic = newGraphic(nil, graphic)

	return f
}

func NewFrameFromImage(path AssetPath, cache, unique bool) (*Frame, error) {
	graphic, err := assets.LoadImage(path, true, false)
	if err != nil {
		return nil, err
	}

	f := new(Frame)
	f._graphic = graphic
	return f, nil
}
