package math

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](v T, vn ...T) (min T) {
	min = v
	for _, v := range vn {
		if min > v {
			min = v
		}
	}
	return min
}

func Max[T constraints.Ordered](v T, vn ...T) (max T) {
	max = v
	for _, v := range vn {
		if max < v {
			max = v
		}
	}
	return max
}

func Clamp[T constraints.Ordered](v, min, max T) T {
	if min > max {
		min, max = max, min
	}
	if v > max {
		return max
	}
	if v < min {
		return min
	}
	return v
}

func Map[T constraints.Integer | constraints.Float | constraints.Complex](v, start, end, scaledStart, scaledEnd T) T {
	return (v - start) / (end - start) * (scaledEnd - scaledStart)
}

// WrapFloat wraps the value "v" between the range "closed" and "opened". The value can wrap to equal the value of "closed" but cannot equal "opened" (it will wrap round to "closed" again).
//
//	WrapInt( 180, -180,  180) // wraps to -180
//	WrapInt(-180, -180,  180) // stays as -180
//	WrapInt(-180,  180, -180) // wraps to 180
//	WrapInt(-180, -180,  180) // stays as 180
func WrapFloat(v, closed, opened float64) float64 {
	if closed == opened {
		panic("min and max cannot be equal")
	}
	min, max := closed, opened
	flipped := min > max
	if flipped {
		min, max = max, min
	}
	if (flipped && v <= min) || (!flipped && v < min) {
		v = max - math.Mod(min-v, max-min)
	}
	v = min + math.Mod(v-min, max-min)
	if v == opened {
		return closed
	}
	return v
}

// WrapInt wraps the value "v" between the range "closed" and "opened". The
// value can wrap to equal the value of "closed" but cannot equal "opened" (it
// will wrap round to "closed" again).
//
//	WrapInt( 180, -180,  180) // wraps to -180
//	WrapInt(-180, -180,  180) // stays as -180
//	WrapInt(-180,  180, -180) // wraps to 180
//	WrapInt(-180, -180,  180) // stays as 180
func WrapInt(v, closed, opened int) int {
	if closed == opened {
		panic("min and max cannot be equal")
	}
	min, max := closed, opened
	flipped := min > max
	if flipped {
		min, max = max, min
	}
	if (flipped && v <= min) || (!flipped && v < min) {
		v = max - (min-v)%(max-min)
	}
	v = min + (v-min)%(max-min)
	if v == opened {
		return closed
	}
	return v
}
