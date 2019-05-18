package main

import "fmt"

func SumOfMultiples(limit int, divisors []int) int {
	totalSum := 0
	for _, divisor := range divisors {
		multiple := divisor
		for multiple < limit {
			totalSum += multiple
			multiple += divisor
		}

	}

	return totalSum
}

func main() {

	fmt.Printf("Sum of multiples is %d\n", SumOfMultiples(12, []int{3, 5}))
	fmt.Printf("Sum of multiples is %d\n", SumOfMultiples(15, []int{3, 4, 5}))
	fmt.Printf("Sum of multiples is %d\n", SumOfMultiples(0, []int{3, 4}))
	fmt.Printf("Sum of multiples is %d\n", SumOfMultiples(5, []int{}))

}
