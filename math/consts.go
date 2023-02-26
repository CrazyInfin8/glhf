package math

import "math"

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
