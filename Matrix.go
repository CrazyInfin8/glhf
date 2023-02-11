package glhf

import "github.com/crazyinfin8/glhf/math"

type Matrix = math.Matrix

func NewMatrix(a, b, c, d, tx, ty float64) Matrix { return math.NewMatrix(a, b, c, d, tx, ty) }

func NewMatrixIdentity() Matrix { return math.NewMatrixIdentity() }

func NewMatrixFromBox(scaleX, scaleY, rotation, tx, ty float64) Matrix {
	return math.NewMatrixFromBox(scaleX, scaleY, rotation, tx, ty)
}
