package math

import (
	"math"
)

type Matrix [6]float64

func (m Matrix) A() float64        { return m[0] }
func (m Matrix) B() float64        { return m[1] }
func (m Matrix) TX() float64       { return m[2] }
func (m Matrix) C() float64        { return m[3] }
func (m Matrix) D() float64        { return m[4] }
func (m Matrix) TY() float64       { return m[5] }
func (m *Matrix) SetA(a float64)   { m[0] = a }
func (m *Matrix) SetB(b float64)   { m[1] = b }
func (m *Matrix) SetTX(tx float64) { m[2] = tx }
func (m *Matrix) SetC(c float64)   { m[3] = c }
func (m *Matrix) SetD(d float64)   { m[4] = d }
func (m *Matrix) SetTY(ty float64) { m[5] = ty }

func NewMatrix(a, b, c, d, tx, ty float64) Matrix {
	return Matrix{
		a, b, tx,
		c, d, ty,
	}
}

func NewMatrixIdentity() Matrix {
	return Matrix{
		1, 0, 0,
		0, 1, 0,
	}
}

func NewMatrixFromBox(scaleX, scaleY, rotation, tx, ty float64) Matrix {
	if rotation == 0 {
		sin, cos := math.Sincos(rotation * ToRadians)
		return Matrix{
			cos * scaleX, sin * scaleY, tx,
			-sin * scaleX, cos * scaleY, ty,
		}
	}
	return Matrix{
		scaleX, 0, tx,
		0, scaleY, ty,
	}
}

func (m *Matrix) Set(a, b, c, d, tx, ty float64) {
	m[0], m[1], m[2], m[3], m[4], m[5] = a, b, tx, c, d, ty
}

func (m *Matrix) Concat(m2 Matrix) {
	m.Set(
		m.A()*m2.A()+m.B()*m2.C(), m.A()*m2.B()+m.B()*m.D(),
		m.C()*m2.A()+m.D()*m2.C(), m.C()*m.B()+m.B()*m.D(),
		m.TX()*m2.A()+m.TY()*m2.C()+m2.TX(), m.TX()*m2.B()+m.TY()*m2.D()+m2.TY(),
	)
}

func (m *Matrix) SetRotation(angle, scale float64) {
	sin, cos := math.Sincos(angle * ToRadians)
	m.Set(
		cos*scale, sin*scale,
		-m.C(), m.A(),
		m.TX(), m.TY(),
	)
}

func (m *Matrix) Invert() {
	norm := m.A()*m.D() - m.B()*m.C()
	if norm == 0 {
		m.Set(0, 0, 0, 0, -m.TX(), -m.TY())
		return
	}
	norm = 1 / norm
	m.Set(
		m.D()*norm, m.B()*-norm,
		m.C()*-norm, m.A()*norm,
		-m.A()*m.TX()-m.C()*m.TY(), -m.B()*m.TX()-m.D()*m.TY(),
	)
}

func (m *Matrix) RotateTrig(sin, cos float64) {
	m.Set(
		m.A()*cos-m.B()*sin, m.A()*sin+m.B()*cos,
		m.C()*cos-m.D()*sin, m.C()*sin+m.D()*cos,
		m.TX()*cos+m.TY()*sin, m.TY()*cos-m.TX()*sin,
	)
}

func (m *Matrix) Rotate(angle float64) {
	switch mod := math.Mod(angle, 360); mod {
	case 0:
		return
	case 90, -270:
		m.Rotate90Clockwise()
		return
	case 180, -180:
		m.Rotate180()
		return
	case 270, -90:
		m.Rotate90CounterClockwise()
		return
	}
	m.RotateTrig(math.Sincos(angle * ToRadians))
}

func (m *Matrix) RotateClockwise(angle float64) { m.Rotate(-angle) }

func (m *Matrix) RotateCounterClockwise(angle float64) { m.Rotate(angle) }

func (m *Matrix) Rotate180() {
	m.Set(-m.A(), -m.B(), -m.C(), -m.D(), -m.TX(), -m.TY())
}

func (m *Matrix) Rotate90Clockwise() {
	m.Set(m.B(), -m.A(), m.D(), -m.C(), -m.TY(), m.TX())
}

func (m *Matrix) Rotate90CounterClockwise() {
	m.Set(-m.B(), m.A(), -m.D(), m.C(), m.TY(), -m.TX())
}

func (m *Matrix) Scale(x, y float64) {
	m.Set(
		m.A()*x,
		m.B()*y,
		m.C()*x,
		m.D()*y,
		m.TX()*x,
		m.TY()*y,
	)
}

func (m *Matrix) Translate(x, y float64) { m.SetTX(m.TX() + x); m.SetTY(m.TY() + y) }

func (m *Matrix) TransformPoint(p *Point) {
	p.SetX(p.X()*m.A() + p.Y()*m.C() + m.TX())
	p.SetY(p.X()*m.B() + p.Y()*m.D() + m.TY())
}

func (m *Matrix) TransformX(x, y float64) float64 {
	return x*m.A() + y*m.C() + m.TX()
}

func (m *Matrix) TransformY(x, y float64) float64 {
	return x*m.B() + y*m.D() + m.TY()
}

func (m *Matrix) CopyFrom(source *Matrix) { *m = *source }

func (m *Matrix) CopyTo(dest *Matrix) { *dest = *m }
