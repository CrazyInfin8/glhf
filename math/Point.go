package math

import (
	"fmt"
	"image"
	"math"
)

type Point struct {
	X, Y float64
}

type Vector = Point

func NewPoint(x, y float64) Point          { return Point{x, y} }
func FromImagePoint(p image.Point) Point   { return Point{float64(p.X), float64(p.Y)} }
func (p *Point) ToImagePoint() image.Point { return image.Point{int(p.X), int(p.Y)} }
func (p *Point) Set(x, y float64)          { p.X, p.Y = x, y }
func (p Point) XY() (x, y float64)         { return p.X, p.Y }
func (p Point) YX() (y, x float64)         { return p.Y, p.X }
func (p *Point) Add(x, y float64)          { p.X += x; p.Y += y }
func (p *Point) AddPoint(p2 Point)         { p.X += p2.X; p.Y += p2.Y }
func (p *Point) Sub(x, y float64)          { p.X -= x; p.Y -= y }
func (p *Point) SubPoint(p2 Point)         { p.X -= p2.X; p.Y -= p2.Y }
func (p *Point) Scale(k float64)           { p.X *= k; p.Y *= k }
func (p *Point) Mult(x, y float64)         { p.X *= x; p.Y *= y }
func (p *Point) MultPoint(p2 Point)        { p.X *= p2.X; p.Y *= p2.Y }
func (p *Point) Div(x, y float64)          { p.X /= x; p.Y /= y }
func (p *Point) DivPoint(p2 Point)         { p.X /= p2.X; p.Y /= p2.Y }

func (p *Point) Rotate(pivot Point, angle float64) {
	sin, cos := math.Sincos(angle * ToRadians)
	dx, dy := p.X-pivot.X, p.Y-pivot.Y
	p.X, p.Y = cos*dx-sin*dy+pivot.X, sin*dx+cos*dy+pivot.Y
}

func (p Point) AngleBetween(p2 *Point) float64 {
	return math.Atan2(p2.Y-p.Y, p2.X-p.X) * ToDegrees
}

func (p Point) IsOrigin() bool {
	return p.X == 0 && p.Y == 0
}

func (p *Point) Transform(mat Matrix) {
	// p.X, p.Y = mat.Apply(p.XY())
	x, y := p.X, p.Y
	p.X = x*mat.A + y*mat.C + mat.TX
	p.Y = x*mat.B + y*mat.D + mat.TY
}
func (p Point) CeilToInts() (x, y int) { return int(math.Ceil(p.X)), int(math.Ceil(p.Y)) }
func (p Point) RoundToInts() (x, y int) { return int(math.Round(p.X)), int(math.Round(p.Y)) }
func (p Point) FloorToInts() (x, y int) { return int(math.Floor(p.X)), int(math.Floor(p.Y)) }
func (p Point) String() string { return fmt.Sprintf("Point{ X: %f, Y: %f }", p.X, p.Y) }

func (p *Point) SetLength(l float64) {
	a := p.Radians()
	p.X, p.Y = math.Sincos(a)
	p.X *= l
	p.Y *= l
}

func (p *Point) SetDegrees(degs float64) {
	l := p.Length()
	p.X, p.Y = math.Sincos(degs * ToRadians)
	p.X *= l
	p.Y *= l
}

func (p *Point) SetRadians(rads float64) {
	l := p.Length()
	p.X, p.Y = math.Sincos(rads)
	p.X *= l
	p.Y *= l
}

func (p Point) DX() float64      { return p.X / p.Length() }
func (p Point) DY() float64      { return p.Y / p.Length() }
func (p Point) Length() float64  { return math.Sqrt(p.X*p.X + p.Y*p.Y) }
func (p Point) Length2() float64 { return p.X*p.X + p.Y*p.Y }
func (p Point) Degrees() float64 { return math.Atan2(p.X, p.Y) * ToDegrees }
func (p Point) Radians() float64 { return math.Atan2(p.X, p.Y) }
func (p Point) IsZero() bool     { return p.X == 0 && p.Y == 0 }
func (p Point) RX() float64      { return -p.Y }
func (p Point) RY() float64      { return p.X }
func (p Point) LX() float64      { return p.Y }
func (p Point) LY() float64      { return -p.X }