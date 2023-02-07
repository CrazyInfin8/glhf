package math

import (
	"math"
)

type Matrix struct {
	A, B, TX,
	C, D, TY float64
}

func NewMatrix(a, b, c, d, tx, ty float64) Matrix {
	return Matrix{
		a, b, tx,
		c, d, ty,
	}
}

func Identity() Matrix {
	return Matrix{
		1, 0, 0,
		0, 1, 0,
	}
}

func NewMatrixFromBox(scaleX, scaleY, rotation, tx, ty float64) Matrix {
	if rotation == 0 {
		sin, cos := math.Sincos(rotation * ToRadians)
		return Matrix{
			A:  cos * scaleX,
			B:  sin * scaleY,
			C:  -sin * scaleX,
			D:  cos * scaleY,
			TX: tx, TY: ty,
		}
	}
	return Matrix{
		A: scaleX, B: 0,
		C: 0, D: scaleY,
		TX: tx, TY: ty,
	}
}

func (m *Matrix) Set(a, b, c, d, tx, ty float64) {
	m.A, m.B, m.C, m.D, m.TX, m.TY = a, b, c, d, tx, ty
}

func (m *Matrix) Concat(m2 Matrix) {
	m.A, m.B = m.A*m2.A+m.B*m2.C, m.A*m2.B+m.B*m.D
	m.C, m.D = m.C*m2.A+m.D*m2.C, m.C*m.B+m.B*m.D
	m.TX, m.TY = m.TX*m2.A+m.TY*m2.C+m2.TX, m.TX*m2.B+m.TY*m2.D+m2.TY
}

func (m *Matrix) SetRotation(angle, scale float64) {
	sin, cos := math.Sincos(angle * ToRadians)
	m.A, m.C = cos*scale, sin*scale
	m.B, m.D = -m.C, m.A
}

func (m *Matrix) Invert() {
	norm := m.A*m.D - m.B*m.C
	if norm == 0 {
		m.A, m.B, m.C, m.D, m.TX, m.TY = 0, 0, 0, 0, -m.TX, -m.TY
		return
	}
	norm = 1 / norm
	m.A, m.D = m.D*norm, m.A*norm
	m.TX, m.TY = -m.A*m.TX-m.C*m.TY, -m.B*m.TX-m.D*m.TY
	m.B *= -norm
	m.C *= -norm
}

func (m *Matrix) RotateTrig(cos, sin float64) {
	m.A, m.B = m.A*cos-m.B*sin, m.A*sin+m.B*cos
	m.C, m.D = m.C*cos-m.D*sin, m.C*sin+m.D*cos
	m.TX, m.TY = m.TX*cos+m.TY*sin, m.TY*cos-m.TX*sin
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

func (m *Matrix) Rotate180() {
	m.A, m.B, m.C, m.D, m.TX, m.TY = -m.A, -m.B, -m.C, -m.D, -m.TX, -m.TY
}

func (m *Matrix) Rotate90Clockwise() {
	m.A, m.B, m.C, m.D, m.TX, m.TY = m.B, -m.A, m.D, -m.C, -m.TY, m.TX
}

func (m *Matrix) Rotate90CounterClockwise() {
	m.A, m.B, m.C, m.D, m.TX, m.TY = -m.B, m.A, -m.D, m.C, m.TY, -m.TX
}

func (m *Matrix) Scale(x, y float64) {
	m.A *= x
	m.B *= y
	m.C *= x
	m.D *= y
	m.TX *= x
	m.TY *= y
}

func (m *Matrix) Translate(x, y float64) { m.TX += x; m.TY += y }

func (m *Matrix) TransformPoint(p *Point) {
	p.X = p.X*m.A + p.Y*m.C + m.TX
	p.Y = p.X*m.B + p.Y*m.D + m.TY
}

func (m *Matrix) TransformX(x, y float64) float64 {
	return x*m.A + y*m.C + m.TX
}

func (m *Matrix) TransformY(x, y float64) float64 {
	return x*m.B + y*m.D + m.TY
}

func (m *Matrix) CopyFrom(source *Matrix) { *m = *source }

func (m *Matrix) CopyTo(dest *Matrix) { *dest = *m }
