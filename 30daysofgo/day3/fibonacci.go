package main

import (
	"fmt"
)

// the basic premise here is that a function closure can hold state
// for any number of variables. here we use a and b to keep track of previous two values
// and always returning new values that are lagging by one. cute trick is to swap a and b
// on the first iteration to give 0. if the series started at 1, then one could assign
// a as 1, b as 0 and return b
// after assigning b, a = a, a+b, but since it starts at 0, we have the following code.
func fibonacci() func() int {
	a := 1
	b := 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

