package math

import (
	"fmt"
	"image"
	"math"
)

type Point [2]float64

type Vector = Point

func NewPoint(x, y float64) Point          { return Point{x, y} }
func FromImagePoint(p image.Point) Point   { return Point{float64(p.X), float64(p.Y)} }
func (p *Point) ToImagePoint() image.Point { return image.Point{int(p[0]), int(p[1])} }
func (p *Point) Set(x, y float64)          { p[0], p[1] = x, y }
func (p Point) X() float64                 { return p[0] }
func (p Point) Y() float64                 { return p[1] }
func (p *Point) SetX(x float64)            { p[0] = x }
func (p *Point) SetY(y float64)            { p[1] = y }
func (p Point) XY() (x, y float64)         { return p[0], p[1] }
func (p Point) YX() (y, x float64)         { return p[1], p[0] }
func (p *Point) SetXY(x, y float64)        { p[0], p[1] = x, y }
func (p *Point) SetYX(y, x float64)        { p[0], p[1] = x, y }
func (p *Point) Add(x, y float64)          { p[0] += x; p[1] += y }
func (p *Point) AddPoint(p2 Point)         { p[0] += p2[0]; p[1] += p2[1] }
func (p *Point) Sub(x, y float64)          { p[0] -= x; p[1] -= y }
func (p *Point) SubPoint(p2 Point)         { p[0] -= p2[0]; p[1] -= p2[1] }
func (p *Point) Scale(k float64)           { p[0] *= k; p[1] *= k }
func (p *Point) Mult(x, y float64)         { p[0] *= x; p[1] *= y }
func (p *Point) MultPoint(p2 Point)        { p[0] *= p2[0]; p[1] *= p2[1] }
func (p *Point) Div(x, y float64)          { p[0] /= x; p[1] /= y }
func (p *Point) DivPoint(p2 Point)         { p[0] /= p2[0]; p[1] /= p2[1] }

func (p *Point) Rotate(pivot Point, angle float64) {
	sin, cos := math.Sincos(angle * ToRadians)
	dx, dy := p[0]-pivot[0], p[1]-pivot[1]
	p[0], p[1] = cos*dx-sin*dy+pivot[0], sin*dx+cos*dy+pivot[1]
}

func (p Point) AngleBetween(p2 *Point) float64 {
	return math.Atan2(p2[1]-p[1], p2[0]-p[0]) * ToDegrees
}

func (p Point) IsOrigin() bool {
	return p[0] == 0 && p[1] == 0
}

func (p *Point) Transform(mat Matrix) {
	p.Set(
		p.X()*mat.A()+p.Y()*mat.C()+mat.TX(),
		p.X()*mat.B()+p.Y()*mat.D()+mat.TY(),
	)
}

func (p *Point) Abs()                   { p.Set(math.Abs(p.X()), math.Abs(p.Y())) }
func (p *Point) Ceil()                  { p.Set(math.Ceil(p.X()), math.Ceil(p.Y())) }
func (p *Point) Round()                 { p.Set(math.Round(p.X()), math.Round(p.Y())) }
func (p *Point) Floor()                 { p.Set(math.Floor(p.X()), math.Floor(p.Y())) }
func (p Point) CeilToInts() (x, y int)  { return int(math.Ceil(p.X())), int(math.Ceil(p.Y())) }
func (p Point) RoundToInts() (x, y int) { return int(math.Round(p.X())), int(math.Round(p.Y())) }
func (p Point) FloorToInts() (x, y int) { return int(math.Floor(p.X())), int(math.Floor(p.Y())) }
func (p Point) String() string          { return fmt.Sprintf("Point{ X: %f, Y: %f }", p.X(), p.Y()) }

func (p *Point) SetLength(l float64) {
	a := p.Radians()
	sin, cos := math.Sincos(a)
	p.Set(cos*l, sin*l)
}

func (p *Point) SetRadians(rads float64) {
	l := p.Length()
	sin, cos := math.Sincos(rads)
	p.Set(cos*l, sin*l)
}

func (p *Point) SetDegrees(degs float64) {
	p.SetRadians(degs * ToRadians)
}

func (p Point) DX() float64      { return p[0] / p.Length() }
func (p Point) DY() float64      { return p[1] / p.Length() }
func (p Point) Length() float64  { return math.Sqrt(p[0]*p[0] + p[1]*p[1]) }
func (p Point) Length2() float64 { return p[0]*p[0] + p[1]*p[1] }
func (p Point) Radians() float64 { return math.Atan2(p[0], p[1]) }
func (p Point) Degrees() float64 { return p.Radians() * ToDegrees }
func (p Point) IsZero() bool     { return p[0] == 0 && p[1] == 0 }
func (p Point) RX() float64      { return -p[1] }
func (p Point) RY() float64      { return p[0] }
func (p Point) LX() float64      { return p[1] }
func (p Point) LY() float64      { return -p[0] }
