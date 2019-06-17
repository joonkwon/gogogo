package main

import "fmt"

type IntSet struct {
	set map[int]bool
}

func (intset *IntSet) add(element int) {
	intset.set[element] = true
}

func (intset *IntSet) getSum() int {
	sum := 0
	for element, _ := range intset.set {
		sum += element
	}
	return sum
}

func SumOfMultiples(limit int, divisors []int) int {

	set := &IntSet{make(map[int]bool)}

	for _, divisor := range divisors {
		multiple := divisor
		for multiple < limit {
			set.add(multiple)
			multiple += divisor
		}

	}

	return set.getSum()
}

func main() {

	fmt.Printf("Sum of multiples is %d\n", SumOfMultiples(12, []int{3, 5}))
	fmt.Printf("Sum of multiples is %d\n", SumOfMultiples(15, []int{3, 4, 5}))
	fmt.Printf("Sum of multiples is %d\n", SumOfMultiples(0, []int{3, 4}))
	fmt.Printf("Sum of multiples is %d\n", SumOfMultiples(5, []int{}))
	fmt.Printf("Sum of multiples is %d\n", SumOfMultiples(20, []int{3, 4, 5, 5, 6}))
}
