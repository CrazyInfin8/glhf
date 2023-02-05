package driver

import (
	"image"
	"image/color"
)

type GraphicsDriver[Img image.Image] graphicsDriver[Img, Image[Img]]

type graphicsDriver[Img image.Image, _Img Image[Img]] interface {
	NewTexture(width, height int) Img
	NewTextureFromImage(img image.Image) Img
}

type Image[Img image.Image] interface {
	image.Image
	SubImage(r image.Rectangle) Img
}

type SubImage struct {
	image.Image
	Rect image.Rectangle
}

func (img SubImage) At(x, y int) color.Color {
	return img.Image.At(img.Rect.Min.X + x, img.Rect.Min.Y + y)
}

func (img SubImage) Bounds() image.Rectangle {
	return img.Rect
}