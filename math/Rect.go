package math

import (
	"image"
	"math"
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

func (r Rect) RotatedBounds(degrees float64, origin Point) Rect {
	degrees = math.Mod(degrees, 360)
	if degrees == 0 {
		return r
	}

	if degrees < 0 {
		degrees += 360
	}

	radians := degrees * ToRadians
	sin, cos := math.Sincos(radians)

	left := -origin.X
	top := -origin.Y
	right := -origin.X + r.Width
	bottom := -origin.Y + r.Height

	if degrees < 90 {
		r.X = r.X + origin.X + cos*left - sin*bottom
		r.Y = r.Y + origin.Y + sin*left + cos*top
	} else if degrees < 180 {
		r.X = r.X + origin.X + cos*right - sin*bottom
		r.Y = r.Y + origin.Y + sin*left + cos*bottom
	} else if degrees < 270 {
		r.X = r.X + origin.X + cos*right - sin*top
		r.Y = r.Y + origin.Y + sin*right + cos*bottom
	} else {
		r.X = r.X + origin.X + cos*left - sin*top
		r.Y = r.Y + origin.Y + sin*right + cos*top
	}

	r.Width, r.Height = math.Abs(cos*r.Height)+math.Abs(sin*r.Width), math.Abs(cos*r.Width)+math.Abs(sin*r.Height)

	return r
}
