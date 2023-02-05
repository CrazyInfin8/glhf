package ebiten

import (
	"image"

	"github.com/crazyinfin8/glhf/driver"
	"github.com/crazyinfin8/glhf/math"
	"github.com/hajimehoshi/ebiten/v2"

	"unsafe"
)

// var _ driver.GraphicProvider = GraphicProvider{}

type GraphicProvider struct{}

func init() {
	driver.Drivers.GraphicProvider = GraphicProvider{}
}

type Graphic struct{ *ebiten.Image }

func (g Graphic) DrawGraphic(src driver.Graphic, matrix math.Matrix) {
	geom := matrixToGeoM(matrix)
	if g, ok := src.(Graphic); ok {
		g.DrawImage(g.Image, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
	} else {
		img := ebiten.NewImageFromImage(src)
		g.DrawImage(img, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
	}
}

func (g Graphic) SubGraphic(r image.Rectangle) driver.Graphic {
	img := g.SubImage(r)
	return Graphic{img.(*ebiten.Image)}
}

// matrixToGeoM converts a [glhf.Matrix] to an [ebiten.GeoM]
func matrixToGeoM(mat math.Matrix) ebiten.GeoM {
	var matrix = struct {
		a_1, b, c, d_1, tx, ty float64
	}{mat.A - 1, mat.B, mat.C, mat.D - 1, mat.TX, mat.TY}

	return *(*ebiten.GeoM)(unsafe.Pointer(&matrix))
}

func (GraphicProvider) NewGraphic(width, height int) driver.Graphic {
	return Graphic{ebiten.NewImage(width, height)}
}

func (GraphicProvider) NewGraphicFromImage(img image.Image) driver.Graphic {
	return Graphic{ebiten.NewImageFromImage(img)}
}
