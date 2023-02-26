package math

import (
	"image"
	"math"
)

type Rect [4]float64

func NewRect(x, y, width, height float64) Rect {
	return Rect{x, y, width, height}
}

func NewRectFromImageRect(r image.Rectangle) Rect {
	return Rect{
		float64(r.Min.X),
		float64(r.Min.Y),
		float64(r.Dx()),
		float64(r.Dy()),
	}
}

func FromOffsetAndSize(offset Point, size Point) Rect {
	return Rect{
		offset.X(),
		offset.Y(),
		size.X(),
		size.Y(),
	}
}

func (r Rect) X() float64                { return r[0] }
func (r Rect) Y() float64                { return r[1] }
func (r Rect) Width() float64            { return r[2] }
func (r Rect) Height() float64           { return r[3] }
func (r *Rect) SetX(x float64)           { r[0] = x }
func (r *Rect) SetY(y float64)           { r[1] = y }
func (r *Rect) SetWidth(w float64)       { r[2] = w }
func (r *Rect) SetHeight(h float64)      { r[3] = h }
func (r Rect) Position() Point           { return Point{r[0], r[1]} }
func (r *Rect) SetPosition(x, y float64) { r[0], r[1] = x, y }
func (r Rect) Size() Point               { return Point{r[2], r[3]} }
func (r *Rect) SetSize(w, h float64)     { r[2], r[3] = w, h }
func (r Rect) Area() float64             { return r[2] * r[3] }

func (r Rect) Top() float64 {
	if r.Height() < 0 {
		return r.Y() + r.Height()
	}
	return r.Y()
}

func (r Rect) Bottom() float64 {
	if r.Width() < 0 {
		return r.Y()
	}
	return r.Y() + r.Height()
}

func (r Rect) Left() float64 {
	if r.Width() < 0 {
		return r[0] + r.Width()
	}
	return r[0]
}

func (r Rect) Right() float64 {
	if r.Width() < 0 {
		return r.X()
	}
	return r.X() + r.Width()
}

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

	left := -origin.X()
	top := -origin.Y()
	right := -origin.X() + r.Width()
	bottom := -origin.Y() + r.Height()

	if degrees < 90 {
		r.SetX(r.X() + origin.X() + cos*left - sin*bottom)
		r.SetY(r.Y() + origin.Y() + sin*left + cos*top)
	} else if degrees < 180 {
		r.SetX(r.X() + origin.X() + cos*right - sin*bottom)
		r.SetY(r.Y() + origin.Y() + sin*left + cos*bottom)
	} else if degrees < 270 {
		r.SetX(r.X() + origin.X() + cos*right - sin*top)
		r.SetY(r.Y() + origin.Y() + sin*right + cos*bottom)
	} else {
		r.SetX(r.X() + origin.X() + cos*left - sin*top)
		r.SetY(r.Y() + origin.Y() + sin*right + cos*top)
	}

	r.SetSize(
		math.Abs(cos*r.Height())+math.Abs(sin*r.Width()),
		math.Abs(cos*r.Width())+math.Abs(sin*r.Height()),
	)

	return r
}
