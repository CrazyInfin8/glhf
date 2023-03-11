package glhf

import (
	"image"
	"image/color"

	"github.com/crazyinfin8/glhf/driver"
)

type (
	Frame struct {
		_graphic
		parent       *Frame
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

func (f *Frame) SubFrame(x, y, width, height int) *Frame {
	if f.UpdateNeeded() {
		f.UpdatePixels()
	}

	texture := f.texture.SubGraphic(image.Rectangle{
		image.Point{X: x, Y: y},
		image.Point{X: x + width, Y: y + height},
	}.Canon())

	subframe := new(Frame)

	subframe._graphic = newGraphic(f._graphic, texture)
	subframe.parent = f

	f.children = append(f.children, subframe._graphic)
	return subframe
}
