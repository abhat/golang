package main

import (
	"golang.org/x/tour/wc"
	"strings"
	//"fmt"
)

func WordCount(s string) map[string]int {
	wcmap := make(map[string]int)
	for _, word := range strings.Fields(s) {
		v, _ := wcmap[word]; 
		wcmap[word] = v + 1 //this works because v will be 0 (default value of the valuetype
	}
	// fmt.Println("Input string is", s)
	// for k, v := range wcmap {
	// 	fmt.Printf("%s is found %d times\n", k, v)
	// }
	return wcmap
}

func main() {
	wc.Test(WordCount)
}

