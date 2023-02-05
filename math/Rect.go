package math

import (
	"image"
)

type Rect struct {
	X, Y, Width, Height float64
}

func NewRect(x, y, width, height float64) Rect {
	return Rect{x, y, width, height}
}

func FromImageRect(r image.Rectangle) Rect {
	return Rect{
		float64(r.Min.X),
		float64(r.Min.Y),
		float64(r.Dx()),
		float64(r.Dy()),
	}
}

func FromOffsetAndSize(offset Point, size Vector) Rect {
	return Rect{
		offset.X,
		offset.Y,
		size.X,
		size.Y,
	}
}

func (r Rect) Top() float64 {
	if r.Height < 0 {
		return r.Y + r.Height
	}
	return r.Y
}

func (r Rect) Bottom() float64 {
	if r.Width < 0 {
		return r.Y
	}
	return r.Y + r.Height
}

func (r Rect) Left() float64 {
	if r.Width < 0 {
		return r.X + r.Width
	}
	return r.X
}

func (r Rect) Right() float64 {
	if r.Width < 0 {
		return r.X
	}
	return r.X + r.Width
}

func (r Rect) Position() Point { return Point{r.X, r.Y} }

func (r *Rect) SetPosition(x, y float64) { r.X, r.Y = x, y }

func (r Rect) Size() Vector { return Vector{r.Width, r.Height} }

func (r *Rect) SetSize(w, h float64) { r.Width, r.Height = w, h }

func (r Rect) Area() float64 { return r.Width * r.Height }

func (r Rect) ToImageRect() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{int(r.Left()), int(r.Top())},
		Max: image.Point{int(r.Right()), int(r.Bottom())},
	}
}
