// Package summultiples implement functions to calculate sum of multiples.
package summultiples

// SumMultiples get limit and divisors and return the sum of multiples of
// the divisors. Duplicated multiples will be added only once.
func SumMultiples(limit int, divs ...int) int {
	multiples := set{
		sum:     0,
		container: make(map[int]bool),
	}
	for _, div := range divs {
		// if divisor is 0, pass
		if div == 0 {
			continue
		}
		// find maximum multiplier
		r := (limit - 1) / div
		// add div * multiplier to multipuls set until r
		for i := 1; i < r+1; i++ {
			multiples.add(i * div)
		}
	}

	return multiples.getSum()
}

// set is a struct for list of unique integers.
type set struct {
	sum     int
	container map[int]bool
}

func (s *set) add(in int) {
	if !s.container[in] {
		s.container[in] = true
		s.sum += in
	}
}

func (s *set) contains(in int) bool {
	_, ok := s.container[in]
	return ok
}

func (s *set) getSum() int {
	return s.sum
}
