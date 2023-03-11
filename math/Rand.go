package math

import (
	"math/rand"
)

func Rand() int {
	return rand.Int()
}

func RandRange(closed, opened int) int {
	var randRange int

	if closed < opened {
		randRange = opened - closed
	} else {
		randRange = closed - opened
	}

	return WrapInt(rand.Intn(randRange), closed, opened)
}
