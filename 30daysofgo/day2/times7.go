package main

import (
	"fmt"
)

func double(x int, iter int) int {
	if iter == 0 {
		return x
	} else {
		return double(x + x, iter-1)
	}
}

func times7(x int) int {
	return double(x, 2) + double(x, 1) + x
}

func times7BitTwiddling(x int) int {
	return x << 2 + x << 1 + x
}
func main () {
	fmt.Println(times7(3))
	fmt.Println(times7BitTwiddling(3))
}
