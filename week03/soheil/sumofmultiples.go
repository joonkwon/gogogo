package main

import (
	"fmt"
)

const testVersion = 1

func SumMultiples(max int, divisors ...int) int {
	result := 0
	for i := 1; i < max; i++ {
		for _, divisor := range divisors {
			if i%divisor == 0 {
				result += i
				break
			}
		}
	}
	return result
}

func SumOfMultiples(limit int, divisors []int) int {
	var result int
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			//fmt.Println("i: ", i)
			//	#fmt.Println("divisor: ", divisor)

			if i%divisor == 0 {
				result += i
				fmt.Println("result: %v i: %v", result, i)
			}

		}
	}

	return result
}

func main() {
	//fmt.Println("Result:", SumOfMultiples(12, []int{3, 5}))
	fmt.Println("Result:", SumOfMultiples(15, []int{3, 4, 5}))
}
