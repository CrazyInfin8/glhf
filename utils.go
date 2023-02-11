package glhf

import (
	"github.com/crazyinfin8/glhf/driver"
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
		panic(name + " is nil. did you forget to store reference?")
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



type Rect = math.Rect

type Graphic = driver.Graphic