package summultiples

import (
	"reflect"
	"testing"
)

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
			t.Errorf("test case failed: %s \n want: %d\n got: %d", c.description, c.expected, actual)
		}
	}
}

func TestAddHappy(t *testing.T) {
	aSet := &set{
		sum:     17,
		container: map[int]bool{
			3: true,
			6: true,
			8: true,
		},
	}
	numberToAdd := 9
	expectedSum := 26
	expectedContain := map[int]bool{
		3: true,
		6: true,
		8: true,
		9: true,
	}

	aSet.add(numberToAdd)

	// if !reflect.DeepEqual(aSet.numbers, expectedNums) {
	// 	t.Fatalf("Add test failed in numbers slice. want: %v, got: %v", expectedNums, aSet.numbers)
	// }
	if aSet.sum != expectedSum {
		t.Fatalf("Add test failed in sum. want: %d, got: %d", expectedSum, aSet.sum)
	}
	if !reflect.DeepEqual(aSet.container, expectedContain) {
		t.Fatalf("Add test failed in contain map. want: %v, got: %v", expectedContain, aSet.container)
	}
}

func BenchmarkSumOfMultiples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			SumMultiples(c.input.limit, c.input.divisors...)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	aSet := &set{
		sum:     17,
		container: map[int]bool{
			3: true,
			6: true,
			8: true,
		},
	}
	for i := 0; i < b.N; i++ {
		aSet.add(i)
	}
}
