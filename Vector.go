package glhf

import "math"

type Vector Point

func NewVector(x, y float64) Vector { return Vector{x, y} }

func (v *Vector) Set(x, y float64) { v.X, v.Y = x, y }

func (v *Vector) DX() float64 { return v.X / v.Length() }

//
func (v *Vector) DY() float64 { return v.Y / v.Length() }

func (v *Vector) Length() float64 { return math.Sqrt(v.X*v.X + v.Y*v.Y) }

func (v *Vector) SetLength(l float64) {
	a := v.Radians()
	v.X, v.Y = math.Sincos(a)
	v.X *= l
	v.Y *= l
}

func (v *Vector) Length2() float64 { return v.X*v.X + v.Y*v.Y }

func (v *Vector) Degrees() float64 { return math.Atan2(v.X, v.Y) * ToDegrees }

func (v *Vector) SetDegrees(degs float64) {
	l := v.Length()
	v.X, v.Y = math.Sincos(degs * ToRadians)
	v.X *= l
	v.Y *= l
}

func (v *Vector) Radians() float64 { return math.Atan2(v.X, v.Y) }

//
func (v *Vector) SetRadians(rads float64) {
	l := v.Length()
	v.X, v.Y = math.Sincos(rads)
	v.X *= l
	v.Y *= l
}

func (v *Vector) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v *Vector) RX() float64 { return -v.Y }

func (v *Vector) RY() float64 { return v.X }

func (v *Vector) LX() float64 { return v.Y }

func (v *Vector) LY() float64 { return -v.X }

func (v *Vector) AsPoint() *Point { return (*Point)(v) }

func (v Vector) ToPoint() Point { return Point(v) }
