package ebiten

import (
	// "GLHF/driver"
	"image"

	"github.com/hajimehoshi/ebiten"
)

// var  _ driver.Driver = {GraphicDriver{}}

type GraphicDriver struct{}

type Image struct{ *ebiten.Image }

func (img Image) SubImage(rect image.Rectangle) Image {
	return Image{img.Image.SubImage(rect).(*ebiten.Image)}
}

func (GraphicDriver) NewTexture(width, height int) Image {
	img, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	return Image{img}
}

func (GraphicDriver) NewTextureFromImage(src image.Image) Image {
	img, _ := ebiten.NewImageFromImage(src, ebiten.FilterDefault)
	return Image{img}
}

func (d GraphicDriver) SubTexture(src Image, rect image.Rectangle) Image {
	img := src.Image.SubImage(rect)
	return Image{img.(*ebiten.Image)}
}


type AnyImage interface {
	image.Image
	SubImage(r image.Rectangle) interface{}
}
