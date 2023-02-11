package ebiten

import (
	"image"
	"image/color"

	"github.com/crazyinfin8/glhf/driver"
	"github.com/crazyinfin8/glhf/math"
	"github.com/hajimehoshi/ebiten/v2"

	"unsafe"
)

type GraphicProvider struct{}

func init() {
	driver.Drivers.GraphicProvider = GraphicProvider{}
}

type Graphic struct{ *ebiten.Image }

func (g Graphic) DrawGraphic(src driver.Graphic, matrix math.Matrix) {
	geom := matrixToGeoM(matrix)
	if src, ok := src.(Graphic); ok {
		g.DrawImage(src.Image, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
		return
	}
	img := ebiten.NewImageFromImage(src)
	g.DrawImage(img, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}

func (g Graphic) SubGraphic(r image.Rectangle) driver.Graphic {
	img := g.SubImage(r)
	return Graphic{img.(*ebiten.Image)}
}

// matrixToGeoM converts a [glhf.Matrix] to an [ebiten.GeoM]
func matrixToGeoM(mat math.Matrix) ebiten.GeoM {
	var matrix = struct {
		a_1, b, c, d_1, tx, ty float64
	}{mat.A() - 1, mat.B(), mat.C(), mat.D() - 1, mat.TX(), mat.TY()}

	return *(*ebiten.GeoM)(unsafe.Pointer(&matrix))
}

func (GraphicProvider) NewGraphic(width, height int) driver.Graphic {
	return Graphic{ebiten.NewImage(width, height)}
}

func (GraphicProvider) NewGraphicFromImage(img image.Image) driver.Graphic {
	return Graphic{ebiten.NewImageFromImage(img)}
}

func (g Graphic) ResizeGraphic(r image.Rectangle) driver.Graphic {
	var geom ebiten.GeoM
	geom.Translate(-float64(r.Min.X), -float64(r.Min.Y))
	img := ebiten.NewImage(r.Dx(), r.Dy())
	img.DrawImage(g.Image, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	return Graphic{img}
}

func (g Graphic) Fill(c color.Color) {
	if c == nil {
		g.Image.Clear()
		return
	}
	g.Image.Fill(c)
}

func (g Graphic) Clone() driver.Graphic {
	img := ebiten.NewImageFromImage(g.Image)
	return Graphic{img}
}