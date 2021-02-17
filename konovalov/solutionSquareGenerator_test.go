package konovalov

import (
	"testing"
)

type squareArguments struct {
	start int
	count int
}

type squaresTestPairs struct {
	values   squareArguments
	expected []int
}

var squareTests = []squaresTestPairs{
	{squareArguments{5, 5}, []int{25, 36, 49, 64, 81}},
	{squareArguments{3, 1}, []int{9}},
	{squareArguments{2, 0}, []int{}},
}

func TestSolutionSquireGenerator(t *testing.T) {
	for _, pairs := range squareTests {
		result := SolutionSquareGenerator(pairs.values.start, pairs.values.count)
		if !IsEqualSlicesOfInt(result, pairs.expected) {
			t.Error(
				"For", pairs.values,
				"expected", pairs.expected,
				"got", result,
			)
		}
	}
}
