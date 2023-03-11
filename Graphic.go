package glhf

import (
	"image"
	"image/color"

	"github.com/crazyinfin8/glhf/driver"
)

type (
	Graphic struct {
		texture driver.Graphic
		pixels  *image.RGBA

		needsUpdate bool
		pixelsRead  bool
		isDestroyed bool

		children []*Graphic
	}

	_graphic   = *Graphic
)

func newGraphic(parent *Graphic, texture driver.Graphic) *Graphic {
	g := new(Graphic)
	g.texture = texture
	return g
}

func (g *Graphic) Width() int  { return g.texture.Bounds().Dx() }
func (g *Graphic) Height() int { return g.texture.Bounds().Dy() }

func (g *Graphic) UpdateNeeded() bool { return g.needsUpdate }
func (g *Graphic) UpdatePixels() {
	if !g.needsUpdate {
		return
	}
	g.needsUpdate = false
	g.texture.ReplacePixels(g.pixels.Pix)
}

func (g *Graphic) LoadPixels() {
	if g.pixelsRead {
		return
	}
	if g.pixels == nil {
		g.pixels = image.NewRGBA(g.texture.Bounds())
	}
	g.texture.LoadPixels(g.pixels.Pix)
	g.pixelsRead = true
}

func (g *Graphic) Set(x, y int, c color.Color) {
	if !g.needsUpdate {
		g.needsUpdate = true
		g.LoadPixels()
	}
	g.pixels.Set(x, y, c)
}

func (g *Graphic) At(x, y int) color.Color {
	if !g.pixelsRead {
		g.LoadPixels()
	}
	return g.pixels.At(x, y)
}

func (g *Graphic) Bounds() image.Rectangle { return g.texture.Bounds() }

func (g *Graphic) Draw(src *Graphic, transform Matrix) {
	g.pixelsRead = false
	g.UpdatePixels()
	src.UpdatePixels()
	g.texture.DrawGraphic(src.texture, transform)

}

func (g *Graphic) Fill(c color.Color) { g.texture.Fill(c) }
