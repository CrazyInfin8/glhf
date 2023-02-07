package glhf

import "github.com/crazyinfin8/glhf/math"
// import m "math"

type Matrix = math.Matrix

func NewMatrix(a, b, c, d, tx, ty float64) Matrix {
	// return Matrix{
	// 	a, b, tx,
	// 	c, d, ty,
	// }
	return math.NewMatrix(a, b, c, d, tx, ty)
}

func Identity() Matrix {
	// return Matrix{
	// 	1, 0, 0,
	// 	0, 1, 0,
	// }
	return math.Identity()
}

// func NewMatrixFromBox(scaleX, scaleY, rotation, tx, ty float64) Matrix {
// 	if rotation == 0 {
// 		sin, cos := m.Sincos(rotation * ToRadians)
// 		return Matrix{
// 			A:  cos * scaleX,
// 			B:  sin * scaleY,
// 			C:  -sin * scaleX,
// 			D:  cos * scaleY,
// 			TX: tx, TY: ty,
// 		}
// 	}
// 	return Matrix{
// 		A: scaleX, B: 0,
// 		C: 0, D: scaleY,
// 		TX: tx, TY: ty,
// 	}
// }
