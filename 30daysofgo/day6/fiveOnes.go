package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Random() {
	if n := rand.Intn(2); n == 0 {
		fmt.Printf("Got zero\n")
		panic(1)
	} else {
		fmt.Printf("Got %d\n", n)
	}
}

func main() {
	fmt.Println("This game continues till we get five consecutive ones with a random throw of a dice")
	for fiveOnes := 0; fiveOnes < 6; fiveOnes++ {
		// sleep for 1s between successive tries.
		time.Sleep(1 * time.Second)
		fmt.Printf("Loop counter is %d\n", fiveOnes)
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("fiveOnes counter got reset at: %d\n", fiveOnes)
					fiveOnes = 0
				} else {
					fmt.Println("defer is a no-op")
				}
			}()
			Random()
		}()
	}
}

