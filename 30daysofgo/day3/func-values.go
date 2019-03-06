package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func times7(fn func(x int) int, x int) int {
	return fn(x) + fn(fn(x)) + x
}


func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	doubler := func(x int) int {
		return x << 1 
	}
	adder5 := func(x int) int {
		return x + 5
	}
	fmt.Println(times7(doubler, 3))
	fmt.Println(times7(adder5, 3))
}

