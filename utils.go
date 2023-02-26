package glhf

import (
	"image"

	"github.com/crazyinfin8/glhf/color"
	"github.com/crazyinfin8/glhf/math"
)

const (
	// ToRadians is a value that, when multiplied by an angle in degrees, will
	// convert that angle to radians
	ToRadians = math.ToRadians

	// ToDegrees is a value that, when multiplied by an angle in radians, will
	// convert that angle to degrees
	ToDegrees = math.ToDegrees

	HalfPi    = math.HalfPi
	QuarterPi = math.QuarterPi
)

func checkNil(ptr interface{}, name string) {
	if ptr == nil {
		panic(name + " is nil. did you forget to create or store reference?")
	}
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}

// unwrap takes the return value and error of a function. It either panics on error, or returns the value.
func unwrap[T interface{}](t T, e error) T {
	if e != nil {
		panic(e)
	}
	return t
}

type Matrix = math.Matrix

func NewMatrix(a, b, c, d, tx, ty float64) Matrix { return math.NewMatrix(a, b, c, d, tx, ty) }
func NewMatrixIdentity() Matrix                   { return math.NewMatrixIdentity() }
func NewMatrixFromBox(scaleX, scaleY, rotation, tx, ty float64) Matrix {
	return math.NewMatrixFromBox(scaleX, scaleY, rotation, tx, ty)
}

type Point = math.Point

func NewPoint(x, y float64) Point                { return math.NewPoint(x, y) }
func NewPointFromImagePoint(p image.Point) Point { return math.NewPointFromImagePoint(p) }

type Rect = math.Rect

func NewRect(x, y, width, height float64) Rect    { return math.NewRect(x, y, width, height) }
func NewRectFromImageRect(r image.Rectangle) Rect { return math.NewRectFromImageRect(r) }

type Color = color.Color

func NewColor(d ...byte) Color                       { return color.NewColor(d...) }
func NewColorFromHex(hex string) Color               { return color.NewColorFromHex(hex) }
func NewColorFromRGB(r, g, b byte) Color             { return color.NewColorFromRGB(r, g, b) }
func NewColorFromRGBA(r, g, b, a byte) Color         { return color.NewColorFromRGBA(r, g, b, a) }
func NewColorFromRGBFloat(r, g, b float64) Color     { return color.NewColorFromRGBFloat(r, g, b) }
func NewColorFromRGBAFloat(r, g, b, a float64) Color { return color.NewColorFromRGBAFloat(r, g, b, a) }
func NewColorFromHSB(hue, saturation, brightness float64) Color {
	return color.NewColorFromHSB(hue, saturation, brightness)
}
func NewColorFromHSBA(hue, saturation, brightness, alpha float64) Color {
	return color.NewColorFromHSB(hue, saturation, brightness)
}
func NewColorFromHSL(hue, saturation, lightness float64) Color {
	return color.NewColorFromHSL(hue, saturation, lightness)
}
func NewColorFromHSLA(hue, saturation, lightness, alpha float64) Color {
	return color.NewColorFromHSLA(hue, saturation, lightness, alpha)
}
