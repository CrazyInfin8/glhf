package driver

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/crazyinfin8/glhf/math"
)

type GraphicOptions struct {
	PixelPerfect bool
	Writeable    bool
}

func DefaultGraphicOptions() GraphicOptions {
	return GraphicOptions{
		PixelPerfect: false,
		Writeable:    false,
	}
}

type GraphicProvider interface {
	NewGraphic(width, height int, opts GraphicOptions) Graphic
	NewGraphicFromImage(img image.Image, opts GraphicOptions) Graphic
	// NewGraphicFromImageWithOptions(img image.Image, opts GraphicOptions) Graphic
}

type Graphic interface {
	draw.Image
	// DrawGraphic draws g onto this graphic with the given translation matrix
	DrawGraphic(src Graphic, matrix math.Matrix)
	SubGraphic(r image.Rectangle) Graphic
	LoadPixels([]byte)
	ReplacePixels([]byte)
	ResizeGraphic(r image.Rectangle) Graphic
	Fill(c color.Color)
	Clone() Graphic
	Destroy()
}
