package summultiples

import "testing"

type testCase struct {
	description string
	input       testInput
	expected    int
}

type testInput struct {
	limit    int
	divisors []int
}

var testCases = []testCase{
	{
		description: "simple case for small numbers",
		input: testInput{
			10,
			[]int{3, 4},
		},
		expected: 30,
	}, {
		description: "simple case for small numbers 2",
		input: testInput{
			18,
			[]int{3, 5},
		},
		expected: 60,
	},
}

func TestSumOfMultiples(t *testing.T) {
	for _, c := range testCases {
		actual := SumMultiples(c.input.limit, c.input.divisors...)
		if actual != c.expected {
			t.Fatalf("test case failed: %s \n want: %d\n got: %d", c.description, c.expected, actual)
		}
	}
}

func BenchmarkSumOfMultiples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			SumMultiples(c.input.limit, c.input.divisors...)
		}
	}
}
