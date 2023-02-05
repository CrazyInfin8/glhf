package main

import (
	glhf "GLHF"
	"math"
)

const (
	// ToRadians is a value that, when multiplied by an angle in degrees, will
	// convert that angle to radians
	ToRadians = math.Pi / 180

	// ToDegrees is a value that, when multiplied by an angle in radians, will
	// convert that angle to degrees
	ToDegrees = 180 / math.Pi

	HalfPi    = math.Pi / 2
	QuarterPi = math.Pi / 4
)

func main() {
	// glhf.NewCamera(0, 0, 0, 0)
	// glhf.Game{}.
}

type PlayState struct {
	*glhf.State
}