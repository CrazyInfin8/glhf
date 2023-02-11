package glhf

import (
	"image"

	"github.com/crazyinfin8/glhf/math"
)

type Point = math.Point

func NewPoint(x, y float64) Point        { return math.NewPoint(x, y) }
func FromImagePoint(p image.Point) Point { return math.FromImagePoint(p) }
