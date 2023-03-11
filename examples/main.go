package main

import "github.com/crazyinfin8/glhf/math"

func main() {
	for i := -10; i < 10; i++ {
		if i != -10 {
			print(", ")
		}
		print("(", i, ":", math.WrapInt(i, -5, -3), ")")
	}
	println()
}
