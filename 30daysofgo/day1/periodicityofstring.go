package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strings"
)

func listProperDivisors(num int) []int {
	divisors := make([]int, 0)
	if num == 1 {
		fmt.Println("(None)")
		return divisors
	}
	for j := 1; j <= num/2; j++ {
		if num%j == 0 {
			divisors = append(divisors, j)
		}
	}
	fmt.Println()
	return divisors
}

func main() {
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input the string which is periodic -> ")
	//inputString, _ := reader.ReadString('\n')
	//inputString = strings.Replace(inputString, "\n", "", -1)
	var inputString string
	fmt.Scanln(&inputString)
	//inputString := "cdcdfdfcdcdabacdcdfdfcdcdabacdcdfdfcdcdabacdcdfdfcdcdaba"
	//inputString := "ababcdcdeababcdcdeababcdcde"
	lenStr := len(inputString)
	foo := listProperDivisors(lenStr)
	fmt.Printf("There are %d divisors for %d. they are %v\n", len(foo), lenStr, foo)

	for _, val := range foo {
		fmt.Printf("Now trying divisor %d for periodicity\n", val)
		count := 0
		for k := 0; k < val; k++ {
			if inputString[k] != inputString[lenStr-val+k] {
				break
			} else {
				count++
			}
		}
		if count == val {
			fmt.Printf("The period of the inputstring is %s\n", inputString[:count])
			return
		}
	}
	fmt.Printf("There is no periodicity - the entire string %s is the period", inputString)
}
